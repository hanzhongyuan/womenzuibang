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
print("分词成功")
ting=[]
a=open('停用词.txt','rb').read()
text=jieba.cut(a)
for i in text:
   ting.append(i)
clearn=[]
for k in title:
    if k not in ting:
      clearn.append(k)
print("处理停用词表")
cipin=Counter(clearn)
for sw in ting:
    del cipin[sw]
print("完成统计词频")
dict=cipin.most_common(10000)
dict_csv=pd.DataFrame(dict)
dict_csv.to_csv('标题词频.csv')




