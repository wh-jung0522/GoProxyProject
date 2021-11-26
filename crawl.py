import requests
import json
import time
KorbitAPIURL = "https://api.korbit.co.kr/v1"

class PulbicAPI():
    def __init__(self) -> None:
        pass

    def korbit_get_ticker(self,currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/ticker?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_detailed_ticker(self,currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/ticker/detailed?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_detailed_ticker_all(self):
        req_url = KorbitAPIURL + '/ticker/detailed/all'
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_orderbook(self,currency='ETH'):
        currency_pair = currency.lower()+"_krw"
        req_url = KorbitAPIURL + '/orderbook?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        print(res_dict)
        return res_dict

    def korbit_get_transactions(self,currency='ETH'):
        currency_pair = currency.lower() + "_krw"
        req_url = KorbitAPIURL + '/transactions?currency_pair={}'.format(currency_pair)
        res = requests.get(req_url)
        res_dict = json.loads(res.text)
        for res_ele in res_dict:
            print(res_ele)
        print(len(res_dict))
        return res_dict


class PrivateAPI():
    def __init__(self,seed_money:int) -> None:
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
                "available":str(seed_money),
                "trade_in_use":"0",
                "withdrawal_in_use" : "0"
            }
        }
        account = {
            "deposit":{},
            "withdrawal": {}
        }
        pass
    ##
    def korbit_get_user_balance():
        pass
    def korbit_get_user_accouts():
        pass
    def korbit_get_user_volume():
        pass
    ##
    def korbit_post_user_buy():
        pass
    def korbit_post_user_sell():
        pass
    def korbit_post_user_cancel():
        pass
    def korbit_get_user_open():
        pass
    def korbit_get_user_orders():
        pass
    def korbit_get_user_transactions():
        pass
    ##
    def korbit_post_user_coin_out():
        pass
    def korbit_post_user_coin_out_cancel():
        pass
    def korbit_get_user_transfer():
        pass
    def korbit_get_user_coin_status():
        pass
    def korbit_post_user_coin_address_assign():
        pass


if __name__ == "__main__":
    pub_api = PulbicAPI()
    res_dict = pub_api.korbit_get_detailed_ticker(currency='ETH')
    print(0)