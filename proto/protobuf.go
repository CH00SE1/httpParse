package proto

/**
 * @title xxx
 * @author xiongshao
 * @date 2022-06-22 15:04:04
 */

import (
	"net"
)

func readMessage(conn net.Conn) {
	/*defer conn.Close()
	buf := make([]byte, 4096, 4096)
	for {
		// 读消息
		cnt, _ := conn.Read(buf)
		player_aaaa := &li5apuu7.Player_aaaa{}
		pData := buf[:cnt]
		// proto解码
		proto.Unmarshal(pData, player_aaaa)
	}*/

}
