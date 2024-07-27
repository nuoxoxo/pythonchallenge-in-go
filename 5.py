from urllib.request import urlopen
import pickle

URL = 'http://www.pythonchallenge.com/pc/def/banner.p'
resp = urlopen(URL)
P = pickle.load(resp)
for p in P:
    line = ''
    for k, v in p:
        line += k * v
    print(line)

print('end/')

