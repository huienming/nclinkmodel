package nclink

var DefaultModel = []byte(`{
	"id": "01",
	"type": "NC_LINK_ROOT",
	"name": "机床模型文件",
	"devices": [
	  {
		"type": "MACHINE",
		"id": "0103",
		"name": "数控机床",
		"description": "数控机床",
		"version": "1.0",
		"configs": [
		  {
			"id": "sample_channel0",
			"type": "SAMPLE_CHANNEL",
			"name": "采样通道",
			"sampleInterval": 2000,
			"uploadInterval": 2000,
			"ids": [
			  {
				"id": "010302"
			  },
			  {
				"id": "010303"
			  },
			  {
				"id": "010305"
			  },
			  {
				"id": "010306"
			  },
			  {
				"id": "010307"
			  },
			  {
				"id": "01035412"
			  }
			]
		  },
		  {
			"id": "sample_channel1",
			"type": "SAMPLE_CHANNEL",
			"name": "采样通道",
			"sampleInterval": 1,
			"uploadInterval": 200,
			"ids": [
			  {
				"id": "0103502001"
			  },
			  {
				"id": "0103502101"
			  },
			  {
				"id": "0103502201"
			  },
			  {
				"id": "0103512001"
			  },
			  {
				"id": "0103512101"
			  },
			  {
				"id": "0103512201"
			  },
			  {
				"id": "0103522001"
			  },
			  {
				"id": "0103522101"
			  },
			  {
				"id": "0103522201"
			  },
			  {
				"id": "0103532002"
			  },
			  {
				"id": "0103532102"
			  },
			  {
				"id": "0103532103"
			  },
			  {
				"id": "01035411"
			  },
			  {
				"id": "01035413"
			  },
			  {
				"id": "01035414"
			  },
			  {
				"id": "010302"
			  }
			]
		  }
		],
		"dataItems": [
		  {
			"id": "010302",
			"name": "机床状态",
			"type": "STATUS"
		  },
		  {
			"id": "010303",
			"name": "进给速度",
			"type": "FEED_SPEED"
		  },
		  {
			"id": "010305",
			"name": "进给倍率",
			"type": "FEED_OVERRIDE"
		  },
		  {
			"id": "010306",
			"name": "主轴倍率",
			"type": "SPDL_OVERRIDE"
		  },
		  {
			"id": "010307",
			"name": "加工件数",
			"type": "PART_COUNT"
		  }
		],
		"components": [
		  {
			"type": "AXIS",
			"number": "0",
			"id": "010350",
			"name": "X轴",
			"description": "",
			"configs": [
			  {
				"id": "01035001",
				"name": "轴名",
				"type": "NAME",
				"value": "X"
			  },
			  {
				"id": "01035002",
				"name": "轴号",
				"type": "NUMBER",
				"value": 0
			  },
			  {
				"id": "01035003",
				"name": "轴类型",
				"type": "TYPE",
				"value": "linear"
			  }
			],
			"components": [
			  {
				"type": "SERVO_DRIVER",
				"id": "01035020",
				"name": "驱动器",
				"description": "",
				"dataItems": [
				  {
					"id": "0103502001",
					"name": "指令位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103502003",
					"name": "指令速度",
					"type": "SPEED"
				  }
				]
			  },
			  {
				"type": "MOTOR",
				"id": "01035021",
				"name": "电机",
				"description": "",
				"dataItems": [
				  {
					"id": "0103502101",
					"name": "负载电流",
					"type": "CURRENT"
				  }
				]
			  },
			  {
				"type": "SCREW",
				"id": "01035022",
				"name": "丝杠",
				"description": "",
				"dataItems": [
				  {
					"id": "0103502201",
					"name": "实际位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103502202",
					"name": "实际速度",
					"type": "SPEED"
				  }
				]
			  }
			]
		  },
		  {
			"type": "AXIS",
			"number": "1",
			"id": "010351",
			"name": "Y轴",
			"description": "",
			"configs": [
			  {
				"id": "01035101",
				"name": "轴名",
				"type": "NAME",
				"value": "Y"
			  },
			  {
				"id": "01035102",
				"name": "轴号",
				"type": "NUMBER",
				"value": 1
			  },
			  {
				"id": "01035103",
				"name": "轴类型",
				"type": "TYPE",
				"value": "linear"
			  }
			],
			"components": [
			  {
				"type": "SERVO_DRIVER",
				"id": "01035120",
				"name": "驱动器",
				"description": "",
				"dataItems": [
				  {
					"id": "0103512001",
					"name": "指令位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103512003",
					"name": "指令速度",
					"type": "SPEED"
				  }
				]
			  },
			  {
				"type": "MOTOR",
				"id": "01035121",
				"name": "电机",
				"description": "",
				"dataItems": [
				  {
					"id": "0103512101",
					"name": "负载电流",
					"type": "CURRENT"
				  }
				]
			  },
			  {
				"type": "SCREW",
				"id": "01035122",
				"name": "丝杠",
				"description": "",
				"dataItems": [
				  {
					"id": "0103512201",
					"name": "实际位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103512202",
					"name": "实际速度",
					"type": "SPEED"
				  }
				]
			  }
			]
		  },
		  {
			"type": "AXIS",
			"number": "2",
			"id": "010352",
			"name": "Z轴",
			"description": "",
			"configs": [
			  {
				"id": "01035201",
				"name": "轴名",
				"type": "NAME",
				"value": "Z"
			  },
			  {
				"id": "01035202",
				"name": "轴号",
				"type": "NUMBER",
				"value": 2
			  },
			  {
				"id": "01035203",
				"name": "轴类型",
				"type": "TYPE",
				"value": "linear"
			  }
			],
			"components": [
			  {
				"type": "SERVO_DRIVER",
				"id": "01035220",
				"name": "驱动器",
				"description": "",
				"dataItems": [
				  {
					"id": "0103522001",
					"name": "指令位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103522003",
					"name": "指令速度",
					"type": "SPEED"
				  }
				]
			  },
			  {
				"type": "MOTOR",
				"id": "01035221",
				"name": "电机",
				"description": "",
				"dataItems": [
				  {
					"id": "0103522101",
					"name": "负载电流",
					"type": "CURRENT"
				  }
				]
			  },
			  {
				"type": "SCREW",
				"id": "01035222",
				"name": "丝杠",
				"description": "",
				"dataItems": [
				  {
					"id": "0103522201",
					"name": "实际位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103522202",
					"name": "实际速度",
					"type": "SPEED"
				  }
				]
			  }
			]
		  },
		  {
			"type": "AXIS",
			"number": "5",
			"id": "010353",
			"name": "C轴",
			"description": "",
			"configs": [
			  {
				"id": "01035301",
				"name": "轴名",
				"type": "NAME",
				"value": "C"
			  },
			  {
				"id": "01035302",
				"name": "轴号",
				"type": "NUMBER",
				"value": 5
			  },
			  {
				"id": "01035303",
				"name": "轴类型",
				"type": "TYPE",
				"value": "rotary"
			  }
			],
			"components": [
			  {
				"type": "SERVO_DRIVER",
				"id": "01035320",
				"name": "C轴驱动器",
				"description": "",
				"dataItems": [
				  {
					"id": "0103532001",
					"name": "指令位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103532002",
					"name": "指令速度",
					"type": "SPEED"
				  }
				]
			  },
			  {
				"type": "MOTOR",
				"id": "01035321",
				"name": "C轴电机",
				"description": "",
				"dataItems": [
				  {
					"id": "0103532101",
					"name": "实际位置",
					"type": "POSITION"
				  },
				  {
					"id": "0103532102",
					"name": "实际速度",
					"type": "SPEED"
				  },
				  {
					"id": "0103532103",
					"name": "负载电流",
					"type": "CURRENT"
				  }
				]
			  }
			]
		  },
		  {
			"type": "CONTROLLER",
			"id": "010354",
			"name": "数控系统",
			"description": "",
			"configs": [
			  {
				"id": "01035404",
				"type": "TOOL_PARAM",
				"name": "刀具参数",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035405",
				"type": "COORDINATE",
				"name": "坐标系",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035406",
				"type": "CONSOLE",
				"name": "指令",
				"setable": true
			  },
			  {
				"id": "01035407",
				"type": "PARAMETER",
				"name": "参数",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035408",
				"type": "FILE",
				"name": "G代码文件",
				"datatype": "HASH",
				"setable": true
			  }
			],
			"dataItems": [
			  {
				"id": "01035409",
				"type": "PROGRAM",
				"name": "主程序名"
			  },
			  {
				"id": "01035410",
				"type": "SUBPROGRAM",
				"name": "子程序名"
			  },
			  {
				"id": "01035411",
				"type": "LINE_NUMBER",
				"name": "指令行号"
			  },
			  {
				"id": "01035412",
				"type": "WARNING",
				"name": "报警"
			  },
			  {
				"id": "01035413",
				"type": "TOOL_NUMBER",
				"name": "刀具号"
			  },
			  {
				"id": "01035414",
				"type": "PROGRAM_NUMBER",
				"name": "程序号"
			  },
			  {
				"id": "01035415",
				"type": "VARIABLE",
				"number": "PROGID_MAP",
				"name": "程序ID映射表"
			  },
			  {
				"id": "01035420",
				"type": "VARIABLE",
				"number": "EVENT",
				"name": "事件"
			  },
			  {
				"id": "01035430",
				"type": "VARIABLE",
				"number": "REG_X",
				"name": "寄存器X",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035431",
				"type": "VARIABLE",
				"number": "REG_Y",
				"name": "寄存器Y",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035432",
				"type": "VARIABLE",
				"number": "REG_F",
				"name": "寄存器F",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035433",
				"type": "VARIABLE",
				"number": "REG_G",
				"name": "寄存器G",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035434",
				"type": "VARIABLE",
				"number": "REG_R",
				"name": "寄存器R",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035435",
				"type": "VARIABLE",
				"number": "REG_W",
				"name": "寄存器W",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035436",
				"type": "VARIABLE",
				"number": "REG_D",
				"name": "寄存器D",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035437",
				"type": "VARIABLE",
				"number": "REG_B",
				"name": "寄存器B",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035438",
				"type": "VARIABLE",
				"number": "REG_P",
				"name": "寄存器P",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035439",
				"type": "VARIABLE",
				"number": "REG_I",
				"name": "寄存器I",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035440",
				"type": "VARIABLE",
				"number": "REG_Q",
				"name": "寄存器Q",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035441",
				"type": "VARIABLE",
				"number": "REG_K",
				"name": "寄存器K",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035442",
				"type": "VARIABLE",
				"number": "REG_T",
				"name": "寄存器T",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035443",
				"type": "VARIABLE",
				"number": "REG_C",
				"name": "寄存器C",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035450",
				"type": "VARIABLE",
				"number": "CHAN_0",
				"name": "通道0数据",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035455",
				"type": "VARIABLE",
				"number": "AXIS_0",
				"name": "轴0数据",
				"datatype": "LIST"
			  },
			  {
				"id": "01035456",
				"type": "VARIABLE",
				"number": "AXIS_1",
				"name": "轴1数据",
				"datatype": "LIST"
			  },
			  {
				"id": "01035457",
				"type": "VARIABLE",
				"number": "AXIS_2",
				"name": "轴2数据",
				"datatype": "LIST"
			  },
			  {
				"id": "01035460",
				"type": "VARIABLE",
				"number": "AXIS_5",
				"name": "轴5数据",
				"datatype": "LIST"
			  },
			  {
				"id": "01035470",
				"type": "VARIABLE",
				"number": "SYS",
				"name": "系统数据",
				"datatype": "LIST"
			  },
			  {
				"id": "01035471",
				"type": "VARIABLE",
				"number": "MACRO",
				"name": "宏变量",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035472",
				"type": "VARIABLE",
				"number": "VAR_AXIS",
				"name": "轴变量",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035478",
				"type": "VARIABLE",
				"number": "VAR_CHAN_0",
				"name": "通道变量",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035482",
				"type": "VARIABLE",
				"number": "VAR_SYS",
				"name": "系统变量",
				"datatype": "LIST",
				"setable": true
			  },
			  {
				"id": "01035483",
				"type": "VARIABLE",
				"number": "VAR_SYSF",
				"name": "浮点系统变量",
				"datatype": "LIST",
				"setable": true
			  }
			]
		  }
		]
	  }
	]
  }`)
