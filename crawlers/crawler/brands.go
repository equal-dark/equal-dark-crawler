package crawler

import (
	. "equal_dark_crawler/crawlers/disturbia"
	. "equal_dark_crawler/crawlers/killstar"
)

var brands = map[int]Crawler{
	1: new(Killstar),
	2: new(Disturbia),
}
