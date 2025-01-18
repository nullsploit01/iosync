import { useLocalSearchParams } from 'expo-router'
import React, { useEffect, useState } from 'react'
import { Text, View } from 'react-native'

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

  return (
    <View>
      <Text>NodeValueScreen</Text>
    </View>
  )
}

export default NodeValueScreen
