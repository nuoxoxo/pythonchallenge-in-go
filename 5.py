from urllib.request import urlopen
import pickle

resp = urlopen('http://www.pythonchallenge.com/pc/def/banner.p')
P = pickle.load(resp)
print('\n'.join(''.join(k * v for k, v in p) for p in P))

