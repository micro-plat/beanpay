# beanpay

创建资金账户,并提供加款、扣费，退款服务; 创建服务包,对服务包使用次数进行增加、扣减、退回操作.

#### 服务列表

<table>
<tr>
<td>分类</td>
<td>服务名</td>
<td>类型</td>
<td>说明</td>
</tr>

<tr>
<td rowspan="6">资金帐户</td>
<td>/account/create </td>
<td rowspan=6>api</td>
<td>添加账户</td>
</tr>

<tr>
<td>/account/balance/add</td>
<td>账户余额加款</td>
</tr>
<tr>
<td>/account/balance/deduct</td>
<td>账户余额扣款</td>
</tr>

<tr>
<td>/account/balance/refund</td>
<td>账户余额退款</td>
</tr>
<tr>
<td>/account/balance/query</td>
<td>账户余额查询</td>
</tr>
<tr>
<td>/account/record/query</td>
<td>账户变动查询</td>
</tr>

<tr>
<td rowspan="6">服务包</td>
<td>/package/create </td>
<td rowspan=6>api</td>
<td>添加服务包</td>
</tr>
<tr>
<td>/package/capacity/add</td>
<td>服务包数量添加</td>
</tr>

<tr>
<td>/package/capacity/deduct</td>
<td>服务包数量扣除</td>
</tr>

<tr>
<td>/package/capacity/refund</td>
<td>服务包数量退回</td>
</tr>

<tr>
<td>/package/capacity/query</td>
<td>服务包数量查询</td>
</tr>

<tr>
<td> /package/record/query </td>
<td>服务包变动查询</td>
</tr>
</table>
