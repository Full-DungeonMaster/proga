import random

random.seed(42)
numbers_A = [random.randint(1, 100) for _ in range(1000)]
numbers_B = [random.randint(1, 100) for _ in range(1000)]

with open('A.txt', 'w') as file_a:
    file_a.write('\n'.join(map(str, numbers_A)))

with open('B.txt', 'w') as file_b:
    file_b.write('\n'.join(map(str, numbers_B)))

with open('A.txt', 'r') as file_a:
    A = list(map(int, file_a.read().splitlines()))

with open('B.txt', 'r') as file_b:
    B = list(map(int, file_b.read().splitlines()))

set_A = set(A)
set_B = set(B)
intersection = set_A & set_B

if not intersection:
    with open('C.txt', 'w') as file_c:
        pass
else:
    count_A = {num: A.count(num) for num in intersection}
    count_B = {num: B.count(num) for num in intersection}
    max_counts = {num: max(count_A[num], count_B[num]) for num in intersection}
    
    sorted_numbers = sorted(intersection)
    result_numbers = []
    
    for num in sorted_numbers:
        result_numbers.extend([num] * max_counts[num])
    
    with open('C.txt', 'w') as file_c:
        file_c.write('\n'.join(map(str, result_numbers)))