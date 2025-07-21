# kbox
A simple cli tool in go/cobra, which can be used to save multiple fyre OCP clusters so that you can switch between them easily.


Major commands

list    - lists all saved cluster names. 
login   - login to a saved cluster. 
save    - to save oc credentials, if logged in you only need to provide the password, else entry for username, password, cluster name will be prompted.  
remove  - remove one or more saved clusters. 

Build the binary

/scripts/build.sh - builds the binary and places it in the scripts dir with name "kbox". 
either run it from here with ./kbox or transfer it to /usr/local/bin to access the utility globally as "kbox"




PS: Feel free to play around the source code and point out improvements :)