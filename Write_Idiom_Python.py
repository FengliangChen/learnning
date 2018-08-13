#!/usr/bin/env python3
# -*- coding: UTF-8 -*-

# Python's way
'''
value exchange
a = 1
b = 2
a,b = b,a
'''

'''
def f(a, b=None):
	if b is None:
		b = []
'''

'''
a = 10
b = 12
c = a if a > b else b
'''

'''
a = '-'
b = ('simon', 'chen', 'age')

print(a.join(b))
print(''.join(b))
'''


'''
name = "Tony"
age = 100
str = "My name is {} and age is {}".format(name, age)
print(str)
'''

#使用列表或者字典comprehension
'''
# bad
mylist = range(20)
odd_list = []
for e in mylist:
    if e % 2 == 1:
        odd_list.append(e)
# good
odd_list = [e for e in mylist if e % 2 == 1]

# bad
user_list = [{'name': 'lucy', 'email': 'lucy@g.com'}, {'name': 'lily', 'email': 'lily@g.com'}]
user_email = {}
for user in user_list:
    if 'email' in user:
        user_email[user['name']] = user['email']
# good
{user['name']: user['email'] for user in user_list if 'email' in user}
'''

#条件判断时，避免直接和True, False, None进行比较(==)
'''
#bad
if l == []:
	pass
#good
if l:
	pass
#bad
if something == None
# good
if something is None
'''


