import csv

input_file = 'anime-2.csv'
output_file = 'anime-2-formatted.csv'
new_id = 1  # Start with ID 1

with open(input_file, 'r', newline='', encoding='utf-8') as infile, \
     open(output_file, 'w', newline='', encoding='utf-8') as outfile:
    
    reader = csv.reader(infile)
    writer = csv.writer(outfile)
    
    # Write the header row if it exists
    header = next(reader, None)
    if header:
        writer.writerow(header)
    
    for row in reader:
        # Replace the first column (ID) with the new sequential ID
        new_row = [str(new_id)] + row[1:]
        writer.writerow(new_row)
        new_id += 1
            
print(f"Formatted CSV saved as {output_file}")
print(f"Total entries processed: {new_id - 1}")
