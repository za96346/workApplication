from flask import Flask, app, request, jsonify
from flask_cors import CORS, cross_origin

app = Flask(__name__)

cors = CORS(app, resources={r"/*": {"origins": "*"}})

@app.route('/', methods=['POST'])
@cross_origin()
def login():
    data = request.get_json()
    account = data['account']
    password = data['password']

    print('original password account ==>',password,account)

    return jsonify('帳號或密碼輸入錯誤'),200

if __name__=="__main__":

    #=accroding_permession_select_route('personal','personalSocialWorker')
    #rint(a)
    # 跨域請求套件
    app.debug = True
    app.run(host='0.0.0.0', port=5001,threaded=True)