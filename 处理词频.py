# coding=gbk
from pandas import Series, DataFrame
import pandas as pd
import numpy as np
import jieba
import jieba.analyse
import sys
import json
from collections import Counter
tran=pd.read_csv('title.csv')
list=pd.DataFrame(tran)
title=[]
for x in list['title']:
    t=jieba.cut(x,cut_all=True)
    for n in t :
        title.append(n)
print("�ִʳɹ�")
ting=[]
a=open('ͣ�ô�.txt','rb').read()
text=jieba.cut(a)
for i in text:
   ting.append(i)
clearn=[]
for k in title:
    if k not in ting:
      clearn.append(k)
print("����ͣ�ôʱ�")
cipin=Counter(clearn)
for sw in ting:
    del cipin[sw]
print("���ͳ�ƴ�Ƶ")
dict=cipin.most_common(10000)
dict_csv=pd.DataFrame(dict)
dict_csv.to_csv('�����Ƶ.csv')




