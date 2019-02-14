def del_ops3(str1, str2):

    # find all common letters in both strings
    common1 = [x for x in str1 if x in str2]
    common2 = [x for x in str2 if x in str1]
    if len(common2) < len(common1):
        common1, common2 = common2, common1

    # find total of strings with 0, 1, or 2 characters, (2 chars - only if c1 != c2)
    if len(common1) == 0 or len(common2) == 0:
        total = len(str1) + len(str2)
    elif (len(common1) == 1 or len(common2) == 1) or (len(common1) == 2 and len(common2) == 2 and common1 != common2):
        total = (len(str1) - 1) + (len(str2) - 1)

    # else, if 2 characters in c1, c2 and c1 != c2 or > 2 characters in c1, c2
    else:

        # create references to c2 indexes of each letter in c1
        refs = defaultdict(list)
        for i, letter in enumerate(common2):
            refs[letter].append(i)

        # find all letters that follow each other (same order) in both strings
        substring = []  # substring == all common letters in same sequence in both strings
        previous = min(refs[common1[0]])
        for i, letter in enumerate(common1):

            # if any c2 index of the current letter in c1 is > the c2 index of previous letter:
                # the current letter follows the previous letter in both c1 and c2
            if any([i > previous for i in refs[letter]]) and all([i != previous for i in refs[letter]]):

                # if the same letter at the same index is not already in substring:
                if all([hash(x) != hash(common2[previous]) for x in substring]):
                    substring.append(common2[previous])

                substring.append(letter)
                previous = min([x for x in refs[letter] if x >= previous])
                # next iteration of previous is always == the smallest index
                # of the current letter that is >= current iteration of previous
                # (always > previous if not first iteration in c1)
                # indexes are never repeated or skipped

            # elif the letter does not follow the same letter in both strings:
                # previous = smallest c2 index of letter that broke sequence/did not follow in both strings
            elif all(refs[letter]) < previous:
                previous = min([x for x in refs[letter]])
            print(i, previous, letter, substring)
        # total == total of all letters - (number of letters in substring * 2)
        total = (len(str1) - len(substring)) + (len(str2) - len(substring))

    return "".join(substring)
    
