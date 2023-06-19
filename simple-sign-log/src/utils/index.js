import jsrsasign from 'jsrsasign';

export default function EncryptData(str) {
    let val = "";
    let rsaPubKey = localStorage.getItem("rsaPubKey")
    let pub = jsrsasign.KEYUTIL.getKey(rsaPubKey)
    for (var i = 0; i < str.length; i++) {
        let jj = encodeURI(str[i])
        if (jj.length == 1) {
            // ascii字符需单独转换
            val += jj.charCodeAt(0).toString(16);
        } else {
            val += jj.split("%").join("");
        }
    }
    let encrpytedData = jsrsasign.KJUR.crypto.Cipher.encrypt(val, pub, 'RSAOAEP256')
    return encrpytedData
}