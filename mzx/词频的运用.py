# coding=gbk
from pandas import Series, DataFrame
import jieba
import pandas as pd
def get_query_recommend (a) :
    tran=pd.read_csv('�����Ƶ.csv')
    list=a
    print("�û����룺"+list)
    jieguo=[]
    jieguo2=[]
    for y in range(1,9999):
        if len(list) <= len(tran['0'][y]):
            if list == tran['0'][y][0:len(list)]:
                jieguo.append(tran['0'][y])

    tran1=pd.read_csv('title.csv')
    list=pd.DataFrame(tran1)
    shuru=a
    print("�û�����:"+shuru)
    for x in range(len(list['title'])):

        if shuru == list['title'][x][0:len(shuru)]:
            jieguo2.append(list['title'][x])
    jieguo2.append(jieguo)
    print("��ʾ��"+str(jieguo2))
    print("�����������ˣ�" + str(jieguo))
a="����"
get_query_recommend(a)