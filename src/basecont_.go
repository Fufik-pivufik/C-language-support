package main

const Version string = "0.002"

const BaseCppFile string = "#include \"../headers/include.hpp\"\n\nint main()\n{\n\tstd::cout << " + `"Hello, World!"` + "<< std::endl;\n\treturn 0;\n}\n"
const BaseIncludeHPP string = "#include <iostream>\n#pragma once\n"
const BaseTestFile string = "\n// In main function here you can write your tests\nint main() {  return 0;  }\n"

const BaseCFile string = "#include \"../headers/include.h\"\n\nint main()\n{\n\tprintf(\"Hello, World!\\n\");\n\treturn 0;\n}\n"
const BaseHFile string = "#include <stdio.h>\n#ifndef HEADER_H\n#define HEADER_H\n\n\n\n#endif\n"
