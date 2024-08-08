# reference: https://en.wikipedia.org/wiki/UTF-8


def str_to_unicode(s):
    ns = ''
    for c in s:
        ns += hex(ord(c)).replace("0x", '\\u')
    return ns


def str_to_binary(s):
    ns = ''
    for c in s:
        ns += bin(ord(c)).replace("0b", "").zfill(16) + " "
    return ns


def str_to_utf8_detail(s):
    for ch in s:
        cbin = bin(ord(ch)).replace("0b", "")
        if cbin.zfill(16) <= bin(int("007F", 16)).zfill(16):
            # 1B
            cbin = cbin.zfill(7)
            a = cbin[:7]
            tmp = "0(%s)" % a
            a0x = hex(int("0" + a, 2)).replace("0x", "\\x")
            print(ch, tmp, "hex:", a0x)
        elif cbin.zfill(16) <= bin(int("07FF", 16)).zfill(16):
            cbin = cbin.zfill(11)
            # 2B
            a, b = cbin[:5], cbin[5:11]
            tmp = "110(%s) 10(%s)" % (a, b)
            a0x = hex(int("110" + a, 2)).replace("0x", "\\x")
            b0x = hex(int("10" + b, 2)).replace("0x", "\\x")
            print(ch, tmp, "hex:", a0x + b0x)
        elif cbin.zfill(16) <= bin(int("FFFF", 16)).zfill(16):
            # 3B
            cbin = cbin.zfill(16)
            a, b, c = cbin[:4], cbin[4:10], cbin[10:16]
            tmp = "1110(%s) 10(%s) 10(%s)" % (a, b, c)
            a0x = hex(int("1110"+a, 2)).replace("0x", "\\x")
            b0x = hex(int("10"+b, 2)).replace("0x", "\\x")
            c0x = hex(int("10"+c, 2)).replace("0x", "\\x")
            print(ch, tmp, "hex:", a0x+b0x+c0x)


if __name__ == '__main__':
    s = "$&"
    print("--------- {} ------------".format(s))
    print("binary", str_to_binary(s))
    print("unicode", str_to_unicode(s))
    print("转化过程")
    str_to_utf8_detail(s)
    print("utf-8", s.encode("utf-8"))

    s = "世界"
    print("--------- {} ------------".format(s))
    print("binary", str_to_binary(s))
    print("unicode", str_to_unicode(s))
    print("转化过程")
    str_to_utf8_detail(s)
    print("utf-8", s.encode("utf-8"))

    # output:
    # --------- $& ------------
    # binary 0000000000100100 0000000000100110 
    # unicode \u24\u26
    # 转化过程
    # $ 0(0100100) hex: \x24
    # & 0(0100110) hex: \x26
    # utf-8 b'$&'
    # --------- 世界 ------------
    # binary 0100111000010110 0111010101001100 
    # unicode \u4e16\u754c
    # 转化过程
    # 世 1110(0100) 10(111000) 10(010110) hex: \xe4\xb8\x96
    # 界 1110(0111) 10(010101) 10(001100) hex: \xe7\x95\x8c
    # utf-8 b'\xe4\xb8\x96\xe7\x95\x8c'


    # 原理讲解
    # binary 0100111000010110 0111010101001100
    # unicode \u4e16\u754c
    # utf-8生成规则：
    # 0000-007F, 7bit, 可表示 2^7=128个字符, 占用1B
    # 0080-07FF, 11bit, 可表示 2^11-2^7=1920个字符，占用2B
    # 0800-FFFF, 16bit, 可表示 2^16-2^11-非法字符=61440个字符，占用3B
    # 10000-10FFFF 占用4B
    # utf-8 b'\xe4\xb8\x96\xe7\x95\x8c'
