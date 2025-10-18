package main

const BaseCppFile string = "#include <iostream>\n\nint main()\n{\n\tstd::cout << " + `"Hello, World!"` + "<< std::endl;\n\treturn 0;\n}\n"
const BaseTestCppFile string = "\n// In main function here you can write your tests\nint main() {  return 0;  }\n"
