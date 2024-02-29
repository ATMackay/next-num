from randgen.randgen import RandomGen

# Print random number sequence
rd = RandomGen()

N = 100

# Print random number sequence
print(f"Generating {N} random numbers!")
for  i in range(N):
    print(rd.next_num())