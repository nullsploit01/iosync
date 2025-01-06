import { useEffect, useState } from 'react'
import { Button, Card, H2, Paragraph, View, XStack } from 'tamagui'

import { nodeAPIService } from '@/src/service/api/node'
import { INode } from '@/src/types/models/node'

const HomeScreen = () => {
  const [_nodes, setNodes] = useState<INode[]>([])

  useEffect(() => {
    fetchNodes()
  }, [])

  const fetchNodes = async () => {
    const response = await nodeAPIService.getNodes()
    setNodes(response.data.data)
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
    <Card elevate size="$4" bordered>
      <Card.Header padded>
        <H2>{node.name}</H2>
        <Paragraph theme="alt2">{node.description}</Paragraph>
      </Card.Header>
      <Card.Footer padded>
        <XStack flex={1} />
        <Button borderRadius="$10">Purchase</Button>
      </Card.Footer>
    </Card>
  )
}

export default HomeScreen
