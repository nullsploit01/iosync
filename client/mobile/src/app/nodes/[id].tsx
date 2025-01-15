import { Stack, useLocalSearchParams } from 'expo-router'
import React, { Fragment, useEffect, useState } from 'react'
import { Card, Text, View, XStack } from 'tamagui'

import { nodeAPIService } from '@/src/service/api/node'
import { INode } from '@/src/types/models/node'
import { parseSearchParamToNumber } from '@/src/utils'

const NodeScreen = () => {
  const { id } = useLocalSearchParams()
  const [_node, setNode] = useState<INode | null>(null)

  useEffect(() => {
    fetchNodeDetails()
  }, [id])

  const fetchNodeDetails = async () => {
    try {
      const parsedId = parseSearchParamToNumber(id)
      const response = await nodeAPIService.getNode(parsedId)
      setNode(response.data.data)
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <Fragment>
      <Stack.Screen options={{ headerTitle: 'Node' }} />
      <View>
        <Card elevate size="$4" margin={16} padding={16}>
          <Card.Header>
            <Text>{_node?.name}</Text>
            <Text>{_node?.description}</Text>
          </Card.Header>
          <XStack>
            <Text color={_node?.is_online ? 'green' : 'red'}>
              {_node?.is_online ? 'Online' : 'Offline'}
            </Text>
          </XStack>
        </Card>
      </View>
    </Fragment>
  )
}

export default NodeScreen
