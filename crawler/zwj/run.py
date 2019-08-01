import requests
import threading
import re, json
import pprint
from utils.docid.decrypt import decrypt_doc_id
from utils.docid.runeval import decrypt_runeval
from utils.document.parse import parse_detail
from utils.vl5x.args import Vjkl5, Vl5x, Number, Guid
from utils.wzws.decrypt import decrypt_wzws

data = {
  "Param": "案件类型:赔偿案件",
  "Index": 2,
  "Page": 10,
  "Order": "法院层级",
  "Direction": "asc",
  "vl5x":None,
  "number":None,
  "guid":None
}


class Spider:
    def __init__(self, proxy="http://127.0.0.1:10809"):
        self.session = requests.Session()
        self.session.headers.update({
            "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36",
        })
        self.proxies = {
          "http": proxy,
          "https": proxy,
        }

    def list_page(self, data=data):
        url = "http://wenshu.court.gov.cn/List/ListContent"
        data["vl5x"] = Vl5x(self.session.cookies.setdefault("vjkl5", Vjkl5()))
        data["number"] = Number()
        data["guid"] = Guid()
        response = self.session.post(url, data=data, proxies=self.proxies)  # 请求1
        text = response.content.decode()

        if "请开启JavaScript并刷新该页" in text:
            dynamic_url = decrypt_wzws(text)
            response = self.session.post(dynamic_url, data=data, proxies = self.proxies)  # 请求2

        json_data = json.loads(response.json())
        #print("列表数据:", json_data)

        runeval = json_data.pop(0)["RunEval"]
        try:
            key = decrypt_runeval(runeval)
        except ValueError as e:
            raise ValueError("返回脏数据") from e
        else:
            print("RunEval解析完成:", key, "\n")

        key = key.encode()
        for item in json_data:
            cipher_text = item["文书ID"]
            print("解密:", cipher_text)
            plain_text = decrypt_doc_id(doc_id=cipher_text, key=key)
            print("成功, 文书ID:", plain_text, "\n")

    def detail_page(self,DocID="f075c337-b647-11e3-84e9-5cf3fc0c2c18"):
        url = "http://wenshu.court.gov.cn/CreateContentJS/CreateContentJS.aspx"
        params = {
            "DocID": DocID,
        }
        response = self.session.get(url, params=params, proxies = self.proxies)  # 请求1
        text = response.content.decode()

        if "请开启JavaScript并刷新该页" in text:
            dynamic_url = decrypt_wzws(text)
            response = self.session.get(dynamic_url, proxies = self.proxies)  # 请求2

        group_dict = parse_detail(response.text)
        pprint(group_dict)


if __name__ == "__main__":
    Spider().list_page()