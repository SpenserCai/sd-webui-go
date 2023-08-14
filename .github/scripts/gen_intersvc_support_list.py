'''
Author: SpenserCai
Date: 2023-08-13 20:20:50
version: 
LastEditors: SpenserCai
LastEditTime: 2023-08-14 17:57:09
Description: file content
'''
import os
import re

header = '''This is the current support list for Intersvc, which is automatically generated through GitHub Action. The list has three fields with the following explanations:

- API Path: The API path in stable-diffusion-webui
- Supported: Whether it supports Intersvc calls
- Checked: Whether it has been verified

{count}
| API Path | Supported | Checked |
| :----: | :----: | :----: |
'''

# 获取intersvc所有inter.go结尾的文件,添加到list，item的内容为inter:inter文件路径，model:model文件路径
def get_all_intersvc_list():
    # 获取_inter.go结尾的文件，通过_inter前面的部分获取model文件
    intersvc_list = []
    for root, dirs, files in os.walk(os.path.join(os.path.dirname(__file__), '../../intersvc')):
        for file in files:
            if file.endswith('_inter.go'):
                intersvc_item = {}
                intersvc_item['inter'] = os.path.join(root, file)
                intersvc_item['model'] = os.path.join(root, file.replace('_inter.go', '_model.go'))
                intersvc_list.append(intersvc_item)
    return intersvc_list

# 通过intersvc_list生成intersvc_support_list.md
def gen_intersvc_support_list():
    # 通过正则匹配inter.go文件中的// API Path: *，获取API Path，通过判断_model.go文件中的type *Response struct {}是否为空，判断是否支持intersvc,通过判断是否有// Checked: True，判断是否已经验证
    intersvc_list = get_all_intersvc_list()
    md_text = ""
    for intersvc in intersvc_list:
        intersvc['path'] = ""
        intersvc['supported'] = ""
        intersvc['checked'] = ""
        with open(intersvc['inter'], 'r') as f:
            for line in f.readlines():
                if line.startswith('// API Path: '):
                    intersvc['path'] = line.replace('// API Path: ', '').strip()
                    break
        with open(intersvc['model'], 'r') as f:
            code = f.read()
            # 找到type *Response struct {\t..\t},如果里面由除了\n和\t以外的字符，说明有内容，说明支持intersvc
            pattern = r'type\s+(\w+)Response\s+struct\s*{([^}]+)}'
            matches = re.findall(pattern, code, re.DOTALL)
            if matches:
                struct_body = matches[0][1].strip()
                if bool(re.sub(r'\s', '', struct_body)):
                    intersvc['supported'] = "√"
            # 如果可以匹配到type *Response = SdApiModel.* 说明支持intersvc
            pattern = r'type\s+(\w+)Response\s*=\s*SdApiModel\.(\w+)'
            matches = re.findall(pattern, code, re.DOTALL)
            if matches:
                intersvc['supported'] = "√"
        if "// Checked: True" in open(intersvc['model']).read():
            intersvc['checked'] = "√"
    # 统计intersvc_list中的intersvc的支持情况
    supported_count = 0
    checked_count = 0
    api_count = len(intersvc_list)
    # 统计的方法为判断intersvc['supported']和intersvc['checked']是否为√，如果是则计数加一
    for intersvc in intersvc_list:
        if intersvc['supported'] == "√":
            supported_count += 1
        if intersvc['checked'] == "√":
            checked_count += 1
    # 生成md文本 用加粗的方式标记统计结果
    count_text = f" **Total API Count: {api_count}** **Supported API Count: {supported_count}**  **Checked API Count: {checked_count}**\n\n"
    for intersvc in intersvc_list:
        md_text += f"| {intersvc['path']} | {intersvc['supported']} | {intersvc['checked']} |\n"
    return md_text,count_text

if __name__ == "__main__":
    t, c = gen_intersvc_support_list()
    print(header.format(count=c)+t)  # 替换为您要更新的内容