# Work Order Lister

Work Order Lister generates a summary text file of Trimble Business Center work order directory outputs. Its purpose is to eliminate the need to search each directory for the desired work order material.

The Work Order Lister program is intended to be run inside a directory whose subdirectories are named beginning with the "GR4_" prefix.

The program will scan these folders for a .txt file and add both the name of the containing directory and the contents of the .txt file to a file named "work_orders.txt" generated in the parent directory.

The "work_orders.txt" file will list all work orders and their descriptions for easy viewing.