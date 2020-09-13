# This script sets up autocompletion for fakegit. To use it, run:
#
# make build/fakegit; source fakegit_completion.sh
#
# From the root of this repo.
#
# In this script, the autocompleter program is set to "./build/fakegit". That
# will work only if you have installed the fakegit executable to the build
# directory, such as if you're at the root level of the fakegit repo and just
# ran "make build/fakegit".
#
# In a "real" release, you would instead pass to complete "-C fakegit" (without
# the "./build/"), because fakegit would be on the PATH.
#
# When we're developing locally, we'd rather not install fakegit to the PATH,
# and so we're fine with an autocompletion script that only works if you have
# the fakegit executable in the current directory.
complete -o bashdefault -o default -C ./build/fakegit fakegit
