# Work Order Lister

Work Order Lister generates a summary text file of Trimble Business Center work order directory outputs. Its purpose is to eliminate the need to search each directory for the desired work order material.

You must create a `prefix.txt` file within the directory that you run the program which lists all of the directory prefixes you would like scanned for text files.

The program will scan these folders for a .txt file and add both the name of the containing directory and the contents of the .txt file to a file named "work_orders.txt" generated in the parent directory.

The "work_orders.txt" file will list all work orders and their descriptions for easy viewing.
