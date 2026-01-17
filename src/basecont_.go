package main

const Version string = "0.002"

const BaseCppFile string = "#include \"../headers/include.hpp\"\n\nint main()\n{\n\tstd::cout << " + `"Hello, World!"` + "<< std::endl;\n\treturn 0;\n}\n"
const BaseIncludeHPP string = "#include <iostream>\n#pragma once\n"
const BaseTestCppFile string = "\n// In main function here you can write your tests\nint main() {  return 0;  }\n"
