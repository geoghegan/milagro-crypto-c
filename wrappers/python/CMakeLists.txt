

function(CURVE_PYTHON_SET tb tf tc nb base nbt m8 mt ct pf)

  if(BUILD_MPIN AND NOT(${pf} STREQUAL "NOT"))

    # Configure mpin wrappers
    file(READ "${PROJECT_SOURCE_DIR}/wrappers/python/mpin_ZZZ.py" tempc)
    string(REGEX REPLACE "@NB@" ${nb} tempc "${tempc}")
    string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
    file(WRITE "${PROJECT_BINARY_DIR}/wrappers/python/mpin_${tc}.py" "${tempc}")

    # Configure mpin test
    file(READ "${PROJECT_SOURCE_DIR}/wrappers/python/TestMPINInstall_ZZZ.py" tempc)
    string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
    file(WRITE "${PROJECT_BINARY_DIR}/wrappers/python/TestMPINInstall_${tc}.py" "${tempc}")

    # Configure mpin time
    file(READ "${PROJECT_SOURCE_DIR}/wrappers/python/TimeMPIN_ZZZ.py" tempc)
    string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
    file(WRITE "${PROJECT_BINARY_DIR}/wrappers/python/TimeMPIN_${tc}.py" "${tempc}")

    # if(${tc} STREQUAL "BN254CX")
      # file(COPY "${PROJECT_SOURCE_DIR}/testVectors/mpin/BN254_CX.json" DESTINATION "${PROJECT_BINARY_DIR}/wrappers/python")
      # file(RENAME "${PROJECT_BINARY_DIR}/wrappers/python/BN254_CX.json" "${PROJECT_BINARY_DIR}/wrappers/python/MPINTestVectors.json")
  	  # file(COPY TestMPIN.py DESTINATION "${PROJECT_BINARY_DIR}/wrappers/python/")
  	  # add_test (test_python_mpin_BN254CX python TestMPIN.py)
    # endif(${tc} STREQUAL "BN254CX")

  	# Install M-PIN module
	  if(${CMAKE_SYSTEM_NAME} MATCHES "Linux" OR ${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
	    INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/python/mpin_${tc}.py DESTINATION "${PYTHON_SITE_PACKAGES}" PERMISSIONS
	            OWNER_WRITE OWNER_READ OWNER_EXECUTE
	            GROUP_READ GROUP_EXECUTE
	            WORLD_READ WORLD_EXECUTE)

	    INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/python/TestMPINInstall_${tc}.py DESTINATION bin PERMISSIONS
	            OWNER_WRITE OWNER_READ OWNER_EXECUTE
	            GROUP_READ GROUP_EXECUTE
	            WORLD_READ WORLD_EXECUTE)
	  endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux" OR ${CMAKE_SYSTEM_NAME} MATCHES "Darwin")

	  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
	    # Windows specific code
	    INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/python/mpin_${tc}.py DESTINATION lib)
	    INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/python/TestMPINInstall_${tc}.py DESTINATION lib)
	  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

	  add_test (test_python_mpin_install_${tc} python TestMPINInstall_${tc}.py)

	endif(BUILD_MPIN AND NOT(${pf} STREQUAL "NOT"))

	if(BUILD_WCC AND NOT(${pf} STREQUAL "NOT"))

	  # Configure mpin wrappers
	  file(READ "${PROJECT_SOURCE_DIR}/wrappers/python/wcc_ZZZ.py" tempc)
	  string(REGEX REPLACE "@NB@" ${nb} tempc "${tempc}")
	  string(REGEX REPLACE "ZZZ" ${tc} tempc "${tempc}")
	  file(WRITE "${PROJECT_BINARY_DIR}/wrappers/python/wcc_${tc}.py" "${tempc}")

	  if(${CMAKE_SYSTEM_NAME} MATCHES "Linux" OR ${CMAKE_SYSTEM_NAME} MATCHES "Darwin")
	    INSTALL(FILES ${PROJECT_BINARY_DIR}/wrappers/python/wcc_${tc}.py DESTINATION "${PYTHON_SITE_PACKAGES}" PERMISSIONS
	            OWNER_WRITE OWNER_READ OWNER_EXECUTE
	            GROUP_READ GROUP_EXECUTE
	            WORLD_READ WORLD_EXECUTE)
	  endif(${CMAKE_SYSTEM_NAME} MATCHES "Linux" OR ${CMAKE_SYSTEM_NAME} MATCHES "Darwin")

	  if(${CMAKE_SYSTEM_NAME} MATCHES "Windows")
	    # Windows specific code
	    INSTALL(FILES ${PROJECT_SOURCE_DIR}/wrappers/python/wcc_${tc}.py DESTINATION lib)
	  endif(${CMAKE_SYSTEM_NAME} MATCHES "Windows")

	endif(BUILD_WCC AND NOT(${pf} STREQUAL "NOT"))

endfunction(CURVE_PYTHON_SET tb tf tc nb base nbt m8 mt ct pf)

string(REGEX REPLACE "," ";" AMCL_CURVE ${AMCL_CURVE})

foreach(CURVE ${AMCL_CURVE})
  if(AMCL_CHUNK STREQUAL "64")
    if(CURVE STREQUAL "ED25519")
        CURVE_PYTHON_SET("256" "25519" "ED25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_PYTHON_SET("256" "25519" "C25519" "32" "56" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_PYTHON_SET("256" "NIST256" "NIST256" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_PYTHON_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_PYTHON_SET("256" "ANSSI" "ANSSI" "32" "56" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_PYTHON_SET("336" "HIFIVE" "HIFIVE" "42" "60" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_PYTHON_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "58" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_PYTHON_SET("384" "NIST384" "NIST384" "48" "56" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_PYTHON_SET("416" "C41417" "C41417" "52" "60" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_PYTHON_SET("528" "NIST521" "NIST521" "66" "60" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_PYTHON_SET("256" "256PMW" "NUMS256W" "32" "56" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_PYTHON_SET("256" "256PME" "NUMS256E" "32" "56" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_PYTHON_SET("384" "384PM" "NUMS384W" "48" "56" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_PYTHON_SET("384" "384PM" "NUMS384E" "48" "56" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_PYTHON_SET("512" "512PM" "NUMS512W" "64" "56" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_PYTHON_SET("512" "512PM" "NUMS512E" "64" "56" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_PYTHON_SET("256" "BN254" "BN254" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_PYTHON_SET("256" "BN254CX" "BN254CX" "32" "56" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_PYTHON_SET("384" "BLS383" "BLS383" "48" "56" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "32")
    if(CURVE STREQUAL "ED25519")
        CURVE_PYTHON_SET("256" "25519" "ED25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_PYTHON_SET("256" "25519" "C25519" "32" "29" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NIST256")
        CURVE_PYTHON_SET("256" "NIST256" "NIST256" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "BRAINPOOL")
        CURVE_PYTHON_SET("256" "BRAINPOOL" "BRAINPOOL" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "ANSSI")
        CURVE_PYTHON_SET("256" "ANSSI" "ANSSI" "32" "28" "256" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "HIFIVE")
        CURVE_PYTHON_SET("336" "HIFIVE" "HIFIVE" "42" "29" "336" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "GOLDILOCKS")
        CURVE_PYTHON_SET("448" "GOLDILOCKS" "GOLDILOCKS" "56" "29" "448" "7" "GENERALISED_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST384")
        CURVE_PYTHON_SET("384" "NIST384" "NIST384" "48" "29" "384" "7" "NOT_SPECIAL" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "C41417")
        CURVE_PYTHON_SET("416" "C41417" "C41417" "52" "29" "414" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NIST521")
        CURVE_PYTHON_SET("528" "NIST521" "NIST521" "66" "28" "521" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256W")
        CURVE_PYTHON_SET("256" "256PMW" "NUMS256W" "32" "28" "256" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_PYTHON_SET("256" "256PME" "NUMS256E" "32" "29" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS384W")
        CURVE_PYTHON_SET("384" "384PM" "NUMS384W" "48" "29" "384" "3" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS384E")
        CURVE_PYTHON_SET("384" "384PM" "NUMS384E" "48" "29" "384" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "NUMS512W")
        CURVE_PYTHON_SET("512" "512PM" "NUMS512W" "64" "29" "512" "7" "PSEUDO_MERSENNE" "WEIERSTRASS" "NOT")
    elseif(CURVE STREQUAL "NUMS512E")
        CURVE_PYTHON_SET("512" "512PM" "NUMS512E" "64" "29" "512" "7" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_PYTHON_SET("256" "BN254" "BN254" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_PYTHON_SET("256" "BN254CX" "BN254CX" "32" "28" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BLS383")
        CURVE_PYTHON_SET("384" "BLS383" "BLS383" "48" "29" "383" "3" "NOT_SPECIAL" "WEIERSTRASS" "BLS")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  elseif(AMCL_CHUNK STREQUAL "16")
    if(CURVE STREQUAL "ED25519")
        CURVE_PYTHON_SET("256" "25519" "ED25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "C25519")
        CURVE_PYTHON_SET("256" "25519" "C25519" "32" "13" "255" "5" "PSEUDO_MERSENNE" "MONTGOMERY" "NOT")
    elseif(CURVE STREQUAL "NUMS256E")
        CURVE_PYTHON_SET("256" "256PME" "NUMS256E" "32" "13" "256" "3" "PSEUDO_MERSENNE" "EDWARDS" "NOT")
    elseif(CURVE STREQUAL "BN254")
        CURVE_PYTHON_SET("256" "BN254" "BN254" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    elseif(CURVE STREQUAL "BN254CX")
        CURVE_PYTHON_SET("256" "BN254CX" "BN254CX" "32" "13" "254" "3" "NOT_SPECIAL" "WEIERSTRASS" "BN")
    else()
      message(FATAL_ERROR "Invalid curve input")
    endif(CURVE STREQUAL "ED25519")
  endif(AMCL_CHUNK STREQUAL "64")
endforeach(CURVE)