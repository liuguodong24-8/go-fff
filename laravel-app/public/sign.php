<?php
// 要签名的 APK X件路经
$apk_path = '14_base.apk';
// 密钥库文件路经
$keystore_path = 'zxcvbnm.jks';
// 秘钥库密码
$keystore_password =  'zxcvbnm';
// 钥别名
$key_alias = 'zxcvbnm';
// 钥密码
$key_password = 'zxcvbnm';
// v1 签名
exec("jarsigner -verbose -sigalg SHAlithRSA -digestalg SHA1 -keystore $keystore_path -storepass $keystore_password -keypass $keystore_password $apk_path $key_alias");
// v2 签名
exec("apksigner sign --ks $keystore_path --ks-pass pass:$keystore_password --key-pass pass:$keystore_password $apk_path");
// v3 签名
exec("apksigner sign --ks $keystore_path --ks-pass pass:$keystore_password --key-pass pass:$keystore_password -v3-signing-enabled true $apk_path");
