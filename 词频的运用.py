# coding=gbk
from pandas import Series, DataFrame
import jieba
import pandas as pd
def get_query_recommend (a) :
    tran=pd.read_csv('标题词频.csv')
    list=a
    print("用户输入："+list)
    jieguo=[]
    jieguo2=[]
    for y in range(1,9999):
        if len(list) <= len(tran['0'][y]):
            if list == tran['0'][y][0:len(list)]:
                jieguo.append(tran['0'][y])

    tran1=pd.read_csv('title.csv')
    list=pd.DataFrame(tran1)
    shuru=a
    print("用户输入:"+shuru)
    for x in range(len(list['title'])):

        if shuru == list['title'][x][0:len(shuru)]:
            jieguo2.append(list['title'][x])
    jieguo2.append(jieguo)
    print("提示："+str(jieguo2))
    print("其他人搜索了：" + str(jieguo))
a="犯罪"
get_query_recommend(a)