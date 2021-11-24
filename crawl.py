import requests
import json
import time
KorbitAPIURL = "https://api.korbit.co.kr/v1"

class PulbicAPI():
    def __init__(self) -> None:
        pass

    def korbit_get_ticker(currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/ticker?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_detailed_ticker(currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/ticker/detailed?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_detailed_ticker_all():
        req_url = KorbitAPIURL + '/ticker/detailed/all'
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_orderbook(currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/orderbook?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_transactions(currency='ETH'):
        currency_pair = currency.lower() + "_krw"
        req_url = KorbitAPIURL + '/transactions?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        for res_ele in res_dict:
            print(res_ele)
        print(len(res_dict))
        return res_dict


class PrivateAPI():
    def __init__(self) -> None:
        korbit_public_API = PulbicAPI()
        coin_wallet_template = {
            "available" : "{1:8f}".format(float(0)),
            "trade_in_use" : "{1:8f}".format(float(0)),
            "withdrawal_in_use" : "{1:8f}".format(float(0)),
            "avg_price": "0",
            "avg_price_updated_at": int(time.time()*1000) # UNIX TIME
        }
        coin_account_template = {
            "address" : ""
        }
        wallet = {
            "krw" : {
                "available":str(int(1e7)),
                "trade_in_use":"0",
                "withdrawal_in_use" : "0"
            }
        }
        account = {
            "deposit":{},
            "withdrawal": {}
        }
        pass

    def korbit_get_user_balance():
        pass