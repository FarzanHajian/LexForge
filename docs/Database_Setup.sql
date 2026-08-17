CREATE USER 'lexforge'@'localhost' IDENTIFIED BY 'lexforge';

CREATE DATABASE `lexforge` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
GRANT ALL ON lexforge.* TO 'lexforge'@'localhost';


CREATE DATABASE `lexforge_unittest` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */ /*!80016 DEFAULT ENCRYPTION='N' */;
GRANT ALL ON lexforge_unittest.* TO 'lexforge'@'localhost';