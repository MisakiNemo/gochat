import JSEncrypt from 'jsencrypt';
import CryptoJS from 'crypto-js';

// 生成AES对称密钥
export const generateAESKey = () => {
    let key = "";
    const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (let i = 0; i < 16; i++)
        key += possible.charAt(Math.floor(Math.random() * possible.length));
    return key;
}

// 使用后端传递过来的公钥加密AES密钥
export const encryptAESKeyWithPublicKey = async (publicKeyPem, aesKeyHex) => {
    const encryptor = new JSEncrypt();
    encryptor.setPublicKey(publicKeyPem);
    return encryptor.encrypt(aesKeyHex);
}

// 使用AES私钥解密
export const decryptAES = async (data, key) => {
    const keyBytes = CryptoJS.enc.Hex.parse(key);
    const decryptor = CryptoJS.AES.decrypt(data, keyBytes, {
        mode: CryptoJS.mode.ECB,
        padding: CryptoJS.pad.NoPadding
    });
    return decryptor.toString(CryptoJS.enc.Utf8);
}

// 使用AES密钥加密
export const encryptAES = (data, key) => {
    const keyBytes = CryptoJS.enc.Hex.parse(key);
    const encrypted = CryptoJS.AES.encrypt(data, keyBytes, {
        mode: CryptoJS.mode.ECB,
        padding: CryptoJS.pad.NoPadding
    });
    return encrypted.toString();
}

export const encryptMD5 = (data) => {
    const encrypted = CryptoJS.MD5(data);
    return encrypted.toString();
}


