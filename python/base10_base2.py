import decimal

D = decimal.Decimal
a = D('2.0')

def dec_to_bin(dec: D) -> str:
    res_int, res_frac = [], []
    dec = str(dec)
    integer, fractional = dec.split('.', 1)
    integer, fractional =int(integer), int(fractional)
    while integer > 0:
        quotient, remainder = divmod(integer, 2)
        integer = quotient
        res_int.append(remainder)
    res_int.reverse()
    while fractional > 0:
       quotient, remainder = divmod(fractional, 2)
       fractional = quotient
       res_frac.append(remainder)
    res_frac.reverse()
    res_int, res_frac = ''.join(map(str, res_int)), ''.join(map(str, res_frac))
    ret = res_int + '.' + res_frac
    return ret

while a <= 5.5:
  print(a, '=binary=>', dec_to_bin(a))
  a+= D('0.5')