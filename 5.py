from urllib.request import urlopen
import pickle

reader = urlopen('http://www.pythonchallenge.com/pc/def/banner.p').read()
print(reader)
print('\n'.join(''.join(k*v for k,v in p) for p in pickle.loads(reader)))

