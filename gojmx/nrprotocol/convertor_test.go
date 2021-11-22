package nrprotocol

// func Test_Convert_Success(t *testing.T) {

// 	testCases := []struct {
// 		name     string
// 		jmxAttr  *JMXAttribute
// 		expected interface{}
// 	}{
// 		{
// 			name: "Double Value",
// 			jmxAttr: &JMXAttribute{
// 				Attribute: "test:type=Cat,name=tomas,attr=FloatValue",
// 				AttributeValue: &JMXAttributeValue{
// 					ValueType:   ValueType_DOUBLE,
// 					DoubleValue: 2.222222,
// 				},
// 			},
// 			expected: float64(2.222222),
// 		},
// 		{
// 			name: "Number Value",
// 			jmxAttr: &JMXAttribute{
// 				Attribute: "test:type=Cat,name=tomas,attr=NumberValue",
// 				AttributeValue: &JMXAttributeValue{
// 					ValueType: ValueType_INT,
// 					IntValue:  3,
// 				},
// 			},
// 			expected: int64(3),
// 		},
// 		{
// 			name: "Bool Value",
// 			jmxAttr: &JMXAttribute{
// 				Attribute: "test:type=Cat,name=tomas,attr=BoolValue",
// 				AttributeValue: &JMXAttributeValue{
// 					ValueType: ValueType_BOOL,
// 					BoolValue: true,
// 				},
// 			},
// 			expected: bool(true),
// 		},
// 		{
// 			name: "Double Value",
// 			jmxAttr: &JMXAttribute{
// 				Attribute: "test:type=Cat,name=tomas,attr=DoubleValue",
// 				AttributeValue: &JMXAttributeValue{
// 					ValueType:   ValueType_DOUBLE,
// 					DoubleValue: 1.2,
// 				},
// 			},
// 			expected: float64(1.2),
// 		},
// 		{
// 			name: "String Value",
// 			jmxAttr: &JMXAttribute{
// 				Attribute: "test:type=Cat,name=tomas,attr=Name",
// 				AttributeValue: &JMXAttributeValue{
// 					ValueType:   ValueType_STRING,
// 					StringValue: "tomas",
// 				},
// 			},
// 			expected: string("tomas"),
// 		},
// 	}

// 	for _, testCase := range testCases {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			assert.Equal(t, testCase.expected, testCase.jmxAttr.GetAttributeValue().GetValue())
// 		})
// 	}
// }
