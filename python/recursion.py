def calc_fatorial(number: int):
    if number == 1:
        return 1
    return number * calc_fatorial(number - 1)
