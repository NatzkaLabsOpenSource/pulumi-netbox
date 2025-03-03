#!/usr/bin/env bash
# WARNING: This file is autogenerated - changes will be overwritten when regenerated by https://github.com/pulumi/ci-mgmt

set -e

original_exec="$0"
original_cmd="$1"

usage() {
  cat <<EOF
NAME
  upstream.sh - Manages applying patches to the upstream submodule.

SYNOPSIS
  ${original_exec} <init|checkout|rebase|check_in|help> [options]

COMMANDS
  init [-f]             Initialize the upstream submodule and applies the
                        patches to the working directory.
  checkout [-f]         Create a branch in the upstream repository with the
                        patches applied as commits.
  rebase [-o] [-i]      Rebase the checked out patches.
  check_in              Write checkedout commits back to patches, add upstream
                        and patches changes to the git staging area and exit
                        checkout mode.
  file_target           Print a file path to depend on in make.
  help                  Print this help message, plus examples.

OPTIONS
  -f   Force the command to run even if the upstream submodule is modified
  -o   The new base commit to rebase the patches on top of
  -i   Run the rebase command interactively
  -h   Print this help message, plus examples
EOF
}

extended_docs() {
  cat <<EOF

DESCRIPTION
  We want to maintain changes to the upstream repository in a way that is easy
  to manage and track. Rather than creating a fork of the upstream repository,
  we maintain a set of patches (in the 'patches' directory) that we can apply
  directly to the upstream code in the 'upstream' submodule. Our patches are
  never pushed to the remote upstream repository.

EXAMPLES
  Discard all changes in upstream and reapply patches to the working directory:
    ${original_exec} init -f

  Moving the patches to a new base commit:
    ${original_exec} checkout
    ${original_exec} rebase -o <new_base_commit>
    ${original_exec} check_in

  Interactively edit the patches:
    ${original_exec} checkout
    ${original_exec} rebase -i
    ${original_exec} check_in

  Add a new patch:
    ${original_exec} checkout
    # Make changes to the upstream repository
    git commit -am "Add new feature"
    ${original_exec} check_in
EOF
}

assert_upstream_exists() {
  if [[ ! -d upstream ]]; then
    echo "No 'upstream' directory detected. Aborting."
    exit 1
  fi
}

assert_not_checked_out() {
  current_branch=$(cd upstream && git --no-pager rev-parse --abbrev-ref HEAD)
  if [[ "${current_branch}" == "pulumi/patch-checkout" ]]; then
    cat <<EOF
Error: 'upstream' submodule checked out on the 'pulumi/patch-checkout' branch.
This was likely caused by running a 'checkout' command but not running
'check_in' afterwards.

To turn the commits in the 'pulumi/patch-checkout' branch back into patches, run:
  ${original_exec} check_in

To disgard changes in the 'pulumi/patch-checkout' branch, use the 'force' flag (-f):
  ${original_exec} ${original_cmd} -f

EOF
    exit 1
  fi

}

assert_is_checked_out() {
  current_branch=$(cd upstream && git --no-pager rev-parse --abbrev-ref HEAD)
  if [[ "${current_branch}" != "pulumi/patch-checkout" ]]; then
    echo "Expected upstream to be checked out on the 'pulumi/patch-checkout' branch, but ${current_branch} is checked out."
    exit 1
  fi
}

assert_no_rebase_in_progress() {
    # Use git to resolve the possible location of files indicating a rebase might be in progress.
  rebase_merge_dir=$(cd upstream && git rev-parse --git-path rebase-merge)
  rebase_apply_dir=$(cd upstream && git rev-parse --git-path rebase-apply)

  if [[ -d "${rebase_merge_dir}" ]] || [[ -d "${rebase_apply_dir}" ]]; then
    echo "rebase still in progress in './upstream'. Please resolve the rebase in"
    exit 1
  fi
}

err_failed_to_apply() {
  cat <<EOF

Failed to apply $1.

Hint: to avoid conflicts when updating the upstream submodule, use the
following commands:

1. '${original_exec} checkout' to create a branch with the patches applied as commits
2. '${original_exec} rebase -o <new_base_commit>' to rebase the patches on top of the
    new upstream commit. Resolve any conflicts and continue the rebase to completion.
3. '${original_exec} check_in' to create an updated set of patches from the commits

Reset the upstream submodule to the previous known good upstream commit before
trying again. This can be done with:

    (cd upstream && git reset --hard <last_known_good_commit>)
    git add upstream

EOF
  exit 1
}

apply_patches() {
  # Iterating over the patches folder in sorted order,
  # apply the patch using a 3-way merge strategy. This mirrors the default behavior of 'git merge'
  cd upstream
  # Allow directory to be empty
  shopt -s nullglob
  for patch in ../patches/*.patch; do
    if ! git apply --3way "${patch}" --allow-empty; then
      err_failed_to_apply "$(basename "${patch}")"
    fi
  done
  shopt -u nullglob
}

clean_rebases() {
  # Clean up any previous in-progress rebases.
  cd upstream
  rebase_merge_dir=$(git rev-parse --git-path rebase-merge)
  rebase_apply_dir=$(git rev-parse --git-path rebase-apply)
  rm -rf "${rebase_merge_dir}"
  rm -rf "${rebase_apply_dir}"
  cd ..
}

clean_branches() {
  cd upstream
  if git show-ref --verify --quiet refs/heads/pulumi/patch-checkout; then
    git branch -D pulumi/patch-checkout
  fi
  if git show-ref --verify --quiet refs/heads/pulumi/checkout-base; then
    git branch -D pulumi/checkout-base
  fi
  if git show-ref --verify --quiet refs/heads/pulumi/original-base; then
    git branch -D pulumi/original-base
  fi
  cd ..
}

init() {
  # Parse additional flags
  while getopts "f" flag; do
    case "${flag}" in
       f) force="true";;
       *) echo "Unexpected option ${flag}"; exit 1;;
    esac
  done

  if [[ ! -d upstream ]]; then
    echo "No 'upstream' directory detected. Skipping init."
    exit 0
  fi

  if [[ "${force}" != "true" ]]; then
    assert_not_checked_out
    assert_no_rebase_in_progress
  fi

  git submodule update --force --init
  cd upstream && git clean -fxd && cd ..

  if [[ "${force}" == "true" ]]; then
    clean_rebases
    clean_branches
  fi
  apply_patches
}

checkout() {
  # Parse additional flags
  while getopts "f" flag; do
    case "${flag}" in
       f) force="true";;
       *) echo "Unexpected option ${flag}"; exit 1;;
    esac
  done

  assert_upstream_exists

  if [[ "${force}" != "true" ]]; then
    assert_not_checked_out
    assert_no_rebase_in_progress
  fi

  git submodule update --force --init
  if [[ "${force}" == "true" ]]; then
    clean_rebases
    clean_branches
  fi

  cd upstream
  git fetch --all

  # Set the 'pulumi/checkout-base' branch to the current commit of the upstream repository
  # This is used to track the base commit of the patches
  # If rebasing, then this must be moved to the new base commit.
  git branch -f pulumi/checkout-base
  # Create a new branch 'pulumi/patch-checkout' which will contain the commits for each patch
  git checkout -B pulumi/patch-checkout

  # Allow directory to be empty
  shopt -s nullglob
  for patch in ../patches/*.patch; do
    if ! git am --3way "${patch}"; then
      err_failed_to_apply "$(basename "${patch}")"
    fi
  done
  shopt -u nullglob

  cat <<EOF

The patches have been checked out as commits in the './upstream' repository.
The 'pulumi/patch-checkout' branch is pointing to the last patch.
The 'pulumi/checkout-base' branch is pointing to the base commit of the patches.

To interactively edit the commits:
  ${original_exec} rebase -i

To change the base of the patches:
  ${original_exec} rebase -o <new_base_commit>

Once you have finished editing the commits, run
  ${original_exec} check_in

EOF
}

rebase() {
  # Parse additional flags
  onto="pulumi/checkout-base"
  interactive="false"
  while getopts "io:" flag; do
    case "${flag}" in
       i) interactive="true";;
       o) onto="${OPTARG}";;
       *) echo "Unexpected option ${flag}"; exit 1;;
    esac
  done

  assert_is_checked_out

  cd upstream
  # Fetch the latest changes from the upstream repository
  git fetch --all
  # Set the "pulumi/original-base" branch to the current base commit of the patches
  git branch -f pulumi/original-base pulumi/checkout-base
  # Set the "pulumi/patch-checkout" branch to track the "pulumi/original-base" branch
  git branch --set-upstream-to=pulumi/original-base pulumi/patch-checkout
  # Set the "pulumi/checkout-base" branch to the new base commit ready for formatting the patches after
  git branch -f pulumi/checkout-base "${onto}"
  # Rebase the 'pulumi/patch-checkout' branch on top of the new base commit
  interactive_flag=""
  if [[ "${interactive}" == "true" ]]; then
    interactive_flag="--interactive"
  fi
  if ! git rebase --onto "${onto}" ${interactive_flag}; then
    echo "Rebase failed. Please resolve the conflicts and run 'git rebase --continue' in the upstream directory. Once the rebase is complete, run '${original_exec} check_in' to write to commits back to patches."
    exit 1
  fi
  cd ..
}

export_patches() {
  # Remove all existing patches before creating the new ones in case they've been renamed or removed.
  rm -f patches/*.patch

  # Extract patches from the commits in the 'pulumi/patch-checkout' branch into the 'patches' directory.
  # Use the 'pulumi/checkout-base' branch to determine the base commit of the patches.
  (cd upstream && git format-patch pulumi/checkout-base -o ../patches --zero-commit --no-signature --no-stat --no-numbered)
}

format_patches() {
  assert_upstream_exists
  assert_is_checked_out
  assert_no_rebase_in_progress

  export_patches
  cat <<EOF
Patches have been created in the 'patches' directory. If you've made changes to
the base, ensure you add 'upstream' to the git stage before running 'init -f'
to exit the checkout mode.
EOF
}

check_in() {
  assert_upstream_exists
  assert_is_checked_out
  assert_no_rebase_in_progress

  export_patches
  # Check out the new base of the patches
  (cd upstream && git checkout pulumi/checkout-base)

  # Add the patches and upstream changes to the git staging area
  git add patches upstream
  # Exit the checkout mode and re-initialize the upstream submodule
  git submodule update --force --init
  apply_patches
  cat <<EOF
Changes to patches and upstream have been staged, exited checkout mode and
re-initializing using updated patches and updated upstream base.
EOF
}

# file_target prints a file path to depend on in make to trigger an init when required.
# Also updates the file timestamp if the submodule needs updating.
file_target() {
  path=.git/modules/upstream/HEAD
  # Don't print a file if it doesn't exist - it's probably not initialized yet.
  if [[ ! -f "${path}" ]]; then
    exit 0
  fi
  # If the submodule is changed, touch the file to trigger a re-init.
  desired_commit=$(git ls-tree HEAD upstream | cut -d ' ' -f3 | cut -f1 || true)
  current_commit=$(cat "${path}")
  if [[ "${desired_commit}" != "${current_commit}" ]]; then
    touch "${path}"
  fi
  echo "${path}"
}

if [[ -z ${original_cmd} ]]; then
  echo "Error: command is required."
  echo
  usage
  extended_docs
  exit 1
fi
# Check for help flag and short-circuit to print usage.
for arg in "$@"; do
  case ${arg} in
    "help"|"-h"|"--help")
      usage
      extended_docs
      exit 0
      ;;
    *)
      ;;
  esac
done

# Remove the command argument from the list of arguments to pass to the command.
shift
case ${original_cmd} in
  init)
    init "$@"
    ;;
  checkout|check_out)
    checkout "$@"
    ;;
  rebase)
    rebase "$@"
    ;;
  format_patches)
    format_patches "$@"
    ;;
  check_in|checkin)
    check_in "$@"
    ;;
  file_target)
    file_target "$@"
    ;;
  *)
    echo "Error: unknown command \"${original_cmd}\"."
    echo
    usage
    exit 1
    ;;
esac
