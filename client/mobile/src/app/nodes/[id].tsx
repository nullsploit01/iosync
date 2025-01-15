import { Stack } from 'expo-router'
import React, { Fragment } from 'react'
import { Text, View } from 'react-native'

const NodeScreen = () => {
  return (
    <Fragment>
      <Stack.Screen options={{ headerTitle: 'Node' }} />
      <View>
        <Text>NodeScreen</Text>
      </View>
    </Fragment>
  )
}

export default NodeScreen
