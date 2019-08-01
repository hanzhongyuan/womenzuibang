import pytesseract
import matplotlib.pyplot as plt
import cv2
import numpy as np
import time

###########################小概率遇验证码，通常返回脏数据

class Img2str():
    def __init__(self,):
        pass
  
    #@staticmethod
    def splitImg(img):      #分割验证码，垂直投影法
        height, width = img.shape[:2]
        vert_proj = []      #
        
        for w in range(width):
            tol = 0
            for h in range(height):
                if img[h][w] == 0 :
                    tol += 1
            vert_proj.append(tol)
        split_point = []
        count = 1
        
        for x in range(len(vert_proj)):     #分割
            if vert_proj[x]==0 and count % 2 == 0:
                split_point.append(x)
                count += 1
            elif vert_proj[x] > 0 and count % 2 != 0:
                split_point.append(x)
                count += 1
        
        if len(split_point) % 2 != 0:
            split_point.append(width)

        if len(split_point) / 2 == 3:       #取中值二次切割，待改
            tag = 0
            lrg = 0
            for x in range(1, 5, 2):
                tmp = split_point[x] -split_point[x-1]
                if lrg < tmp:
                    tag = x
                    lrg = tmp
            mid = int((split_point[tag] + split_point[tag-1])/2)
            split_point.insert(tag, mid)
            split_point.insert(tag, mid)
        elif len(split_point) / 2 != 4:
            return -1
        
        return split_point


    def Test(filename):

        img = preprocImg1(filename)     #图像预处理
        cv2.imwrite('img/proc/p.png', img)
        
        splitL = Img2str.splitImg(img)
        if splitL != -1:
            for x in range(1, len(splitL), 2):
                cv2.imwrite('img/proc/p{}.png'.format(int((x+1)/2)), img[0:-1,splitL[x-1]:splitL[x]])
        else:
            print('验证码分割失败\n')
            return -1
        
        '''for x in range(1,4):
            print(pytesseract.image_to_string('img/proc/p{}.png'.format(x)))'''

    def img2str(self, img):
        pass


def preprocImg1(filename):       #用于去噪点和细横线
    img = cv2.imread(filename)      #BGR顺序读取
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)             #转换成单通道灰度图 -> img.shape
    ret, img = cv2.threshold(img, 0, 255,  cv2.THRESH_OTSU  | cv2.THRESH_BINARY)         #二值化, OTSU大津法
    #img = cv2.resize(img, (img.shape[1],img.shape[0]), (2,2),interpolation=cv2.INTER_LINEAR)      #插值
    
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, (1,2))
    img = cv2.morphologyEx(img, cv2.MORPH_CLOSE, kernel)   #闭运算
    kernel = cv2.getStructuringElement(cv2.MORPH_RECT, (2,2))        
    img = cv2.morphologyEx(img, cv2.MORPH_CLOSE, kernel)
    return img


'''def preprocImg2(filename):
    img = cv2.imread(filename)      #BGR顺序读取
    for x in range(img.shape[0]):
        for y in range(img.shape[1]):
            if img[x][y][0] <=60 and img[x][y][1] <=60 and img[x][y][2] <=60 and \
            img[x][y][0] - img[x][y][1] <= 10 and img[x][y][2] - img[x][y][1] <= 10 \
            and img[x][y][0] - img[x][y][1] <= 10:
                img[x][y][:] = 255
    img = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)             #转换成单通道灰度图 -> img.shape
    ret, img = cv2.threshold(img, 0, 255,  cv2.THRESH_OTSU  | cv2.THRESH_BINARY)         #二值化, OTSU大津法
    return img'''


if __name__ == "__main__":
    Img2str.Test('img/1.png')