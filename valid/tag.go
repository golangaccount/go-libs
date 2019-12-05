package valid

/*
* 需要使用valid函数进行验证的数据只能是struct
* 通过struct中的field的tag信息来进行数据处理
* [] :用来进行包裹验证的参数信息
* ;  :用于分割验证规则
* ： :用作转移字符使用
* ,  :用于进行参数中数据分割
* ...:用于进行range判断
* 如果参数中使用了特殊的字符，请使用对应的转义字符进行转义或使用base64进行编码
* `valid:"validtype[parm];validtype[parm]",valid`
* 验证规则必须优先于类型注册
 */
