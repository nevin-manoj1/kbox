# kbox
A simple cli tool in go/cobra, which can used to save multiple fyre OCP clusters so that you can swicth between them easily.

Major commands

save    - to save oc credentials
login   - login to a saved cluster
list    - lists all saved clsuter names

Obtain the binary
/scripts/build.sh - builds the binary and places it in the scripts dir with name "kbox"
either run it from here with ./kbox or transfer it to /usr/local/bin to access the utility globally as "kbox"




Feel free to play around the source code and point out improvements :)