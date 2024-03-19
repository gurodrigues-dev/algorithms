def quicksort(array: list) -> list:

    if len(array) < 2:
        return array
    
    
    pivo = array[0]

    menores = [i for i in array[1:] if i <= pivo]
    maiores = [i for i in array[1:] if i > pivo]

    return quicksort(menores) + [pivo] + quicksort(maiores)

print(quicksort([10, 5, 2, 3, 8, 4, 9]))