import csv
from collections import defaultdict

input_file = 'anime-2.csv'
output_file = 'anime-2-formatted.csv'
start_id = 3804

id_count = defaultdict(int)
duplicate_ids = set()

# First pass: count occurrences of each ID
with open(input_file, 'r', newline='', encoding='utf-8') as infile:
    reader = csv.reader(infile)
    next(reader, None)  # Skip header
    for row in reader:
        try:
            id_value = int(row[0])
            id_count[id_value] += 1
            if id_count[id_value] > 1:
                duplicate_ids.add(id_value)
        except ValueError:
            pass

# Second pass: process and write to output file
with open(input_file, 'r', newline='', encoding='utf-8') as infile, \
     open(output_file, 'w', newline='', encoding='utf-8') as outfile:
    
    reader = csv.reader(infile)
    writer = csv.writer(outfile)
    
    header = next(reader, None)
    if header:
        writer.writerow(header)
    
    for row in reader:
        try:
            id_value = int(row[0])
            if id_value in duplicate_ids:
                print(f"Duplicate ID found: {id_value}")
            writer.writerow(row)
        except ValueError:
            new_row = [str(start_id)] + row
            writer.writerow(new_row)
            start_id += 1

print(f"Formatted CSV saved as {output_file}")
print(f"Total number of duplicate IDs: {len(duplicate_ids)}")
print(f"Duplicate IDs: {sorted(duplicate_ids)}")
