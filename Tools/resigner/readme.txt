安卓包重签步骤:
		1.把安卓apk和keystore签名文件放到package目录下
		2.配置config.json下的字段
				(1) package_name --> apk文件名,不包含.apk后缀
				(2) keystore --> apk签名文件名,包含后缀
				(3) alias --> apk签名文件中的alias
				(4) alias_password --> apk签名文件中alias对应的密码
		3.配置input目录下的重签版本,例如:
			input --
				| -- 10054
						| -- apk_config.json (对应在重签包中要删除的文件或文件夹)
						| -- apk (对应apk包下的文件结构,把你要新增的文件或文件夹按对应apk目录创建)
		4.双击resigner_android.bat
		5.重签后的所有apk包都将放到output文件夹中

苹果包重签步骤:
		1.请在mac系统中安装ruby和sigh
		2.把苹果ipa和证书xxx.mobileprovision(必须重命名为embedded.mobileprovison)放到package目录下
		3.配置config.json下的字段
				(1) package_name --> ipa文件名,不包含.ipa后缀
				(2) ios_signing_identity --> 开发者证书对应的Signing Identity
		4.配置input目录下的重签版本,例如:
			input --
				| -- 10054
						| -- ipa_config.json (对应在重签包中要删除的文件或文件夹)
						| -- ipa (对应ipa包下的文件结构,把你要新增的文件或文件夹按对应ipa目录创建)
		5.终端运行 sh resigner_ios.sh
		6.重签后的所有ipa包都将放到output文件夹中