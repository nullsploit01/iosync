import { router } from 'expo-router'
import { useEffect, useState } from 'react'
import { TouchableOpacity } from 'react-native'
import { Button, Card, H2, H4, Paragraph, View, XStack } from 'tamagui'

import { nodeAPIService } from '@/src/service/api/node'
import { INode } from '@/src/types/models/node'

const HomeScreen = () => {
  const [_nodes, setNodes] = useState<INode[]>([])

  useEffect(() => {
    fetchNodes()
  }, [])

  const fetchNodes = async () => {
    try {
      const response = await nodeAPIService.getNodes()
      setNodes(response.data.data)
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <View>
      {_nodes.map((node) => (
        <NodeCard key={node.id} node={node} />
      ))}
    </View>
  )
}

const NodeCard = ({ node }: { node: INode }) => {
  return (
    <TouchableOpacity onPress={() => router.push(`/nodes/${node.id}`)}>
      <Card elevate size="$4" bordered>
        <Card.Header padded>
          <H4>#{node?.id}</H4>
          <H2>{node.name}</H2>
          <Paragraph theme="alt2">{node.description}</Paragraph>
        </Card.Header>
        <Card.Footer padded>
          <XStack flex={1} />
          <Button
            borderRadius="$10"
            backgroundColor={`${node.is_active ? '$green9' : '$gray8'}`}
          ></Button>
        </Card.Footer>
      </Card>
    </TouchableOpacity>
  )
}

export default HomeScreen
