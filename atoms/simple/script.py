import sys

# Check if the user has provided exactly two arguments
if len(sys.argv) != 3:
    print("Usage: python hello.py <arg1> <arg2>")
    sys.exit(1)

# Extract the arguments
arg1 = sys.argv[1]
arg2 = sys.argv[2]

# Print the Hello World message with the arguments
print(f"Hello, World! {arg1} {arg2}")
