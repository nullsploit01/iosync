import { useLocalSearchParams } from 'expo-router'
import React, { useEffect, useState } from 'react'
import { Dimensions, Text, View } from 'react-native'
import { LineChart } from 'react-native-chart-kit'

import { nodeAPIService } from '@/src/service/api/node'
import { INodeAPIValue } from '@/src/types/models/node'
import { parseSearchParamToString } from '@/src/utils'

const NodeValueScreen = () => {
  const { api_key } = useLocalSearchParams()
  const [_nodeApiKeyValues, setNodeApiKeyValues] = useState<INodeAPIValue[]>([])

  useEffect(() => {
    fetchNodeValues()
  }, [api_key])

  const fetchNodeValues = async () => {
    try {
      const parsedApiKey = parseSearchParamToString(api_key)
      const response = await nodeAPIService.getNodeValuesByApiKey(parsedApiKey)
      setNodeApiKeyValues(response.data.data)
    } catch (error) {
      console.log(error)
    }
  }

  const screenWidth = Dimensions.get('window').width
  const chartConfig = {
    backgroundGradientFrom: '#1E2923',
    backgroundGradientFromOpacity: 0,
    backgroundGradientTo: '#08130D',
    backgroundGradientToOpacity: 0.5,
    color: (opacity = 1) => `rgba(26, 255, 146, ${opacity})`,
    strokeWidth: 2, // optional, default 3
    barPercentage: 0.5,
    useShadowColorFromDataset: false // optional
  }

  const data = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June'],
    datasets: [
      {
        data: [20, 45, 28, 80, 99, 43],
        color: (opacity = 1) => `rgba(134, 65, 244, ${opacity})`, // optional
        strokeWidth: 2 // optional
      }
    ],
    legend: ['Rainy Days'] // optional
  }

  return (
    <View>
      <Text>NodeValueScreen</Text>
      <LineChart data={data} width={screenWidth} height={220} chartConfig={chartConfig} />
    </View>
  )
}

export default NodeValueScreen
