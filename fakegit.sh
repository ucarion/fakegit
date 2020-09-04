# This script sets up autocompletion for fakegit. To use it, run:
#
# make fakegit; source fakegit.sh
#
# From the root of this repo.
#
# In this script, the autocompleter program is set to "./fakegit". That will
# work only if you have installed the fakegit executable to the current working
# directory, such as if you're at the root level of the fakegit repo and just
# ran "make fakegit".
#
# In a "real" release, you would instead pass to complete "-C fakegit" (without
# the "./"), because fakegit would be on the PATH.
#
# When we're developing locally, we'd rather not install fakegit to the PATH,
# and so we're fine with an autocompletion script that only works if you have
# the fakegit executable in the current directory.
complete -o bashdefault -o default -C ./fakegit fakegit
