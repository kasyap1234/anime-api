import csv

input_file = 'anime-2.csv'
output_file = 'anime-2-formatted.csv'
start_id = 3804  

with open(input_file, 'r', newline='', encoding='utf-8') as infile, \
     open(output_file, 'w', newline='', encoding='utf-8') as outfile:
    
    reader = csv.reader(infile)
    writer = csv.writer(outfile)
    
    # Write the header row if it exists
    header = next(reader, None)
    if header:
        writer.writerow(header)
    
    for row in reader:
        try:
            # Check if the first field can be converted to an integer
            int(row[0])
            # If it can, it's likely an ID, so write the row as is
            writer.writerow(row)
        except ValueError:
            # If it can't, it's likely missing an ID, so add one
            new_row = [str(start_id)] + row
            writer.writerow(new_row)
            start_id += 1

print(f"Formatted CSV saved as {output_file}")
