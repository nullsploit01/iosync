import { ChevronRight } from '@tamagui/lucide-icons'
import { Stack, useLocalSearchParams, useRouter } from 'expo-router'
import React, { Fragment, useEffect, useState } from 'react'
import { Card, ListItem, Text, XStack, YGroup, YStack } from 'tamagui'

import { nodeAPIService } from '@/src/service/api/node'
import { INodeWithAPIKeys } from '@/src/types/models/node'
import { parseSearchParamToNumber } from '@/src/utils'

const NodeScreen = () => {
  const router = useRouter()
  const { id } = useLocalSearchParams()
  const [_node, setNode] = useState<INodeWithAPIKeys | null>(null)

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
      <YStack>
        <Card elevate size="$4" margin={16} padding={16}>
          <Card.Header>
            <Text>#{_node?.id}</Text>
            <Text>{_node?.name}</Text>
            <Text>{_node?.description}</Text>
          </Card.Header>
          <XStack justifyContent="flex-end">
            <Text fontSize="$6" color={_node?.is_online ? 'green' : 'red'}>
              {_node?.is_online ? 'Online' : 'Offline'}
            </Text>
          </XStack>
        </Card>
      </YStack>
      <YGroup alignSelf="center" bordered width={240} size="$5">
        {_node &&
          _node.api_keys &&
          _node.api_keys.map((x) => {
            return (
              <YGroup.Item key={x.id}>
                <ListItem
                  hoverTheme
                  pressTheme
                  onPress={() => router.navigate(`/nodes/${x.id}/values/${x.api_key}`)}
                  title={x.description}
                  subTitle={x.api_key}
                  iconAfter={ChevronRight}
                />
              </YGroup.Item>
            )
          })}
      </YGroup>
    </Fragment>
  )
}

export default NodeScreen
