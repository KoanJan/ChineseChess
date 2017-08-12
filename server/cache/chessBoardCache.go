package cache

import (
	"errors"
	"sync"

	"ChineseChess/server/models"
)

var boardCache *chessBoardCache

type chessBoardCache struct {
	cache     map[string]*models.ChessBoard
	cachelock *chessBoardCacheLock
}

type chessBoardCacheLock struct {
	locks map[string]*sync.Mutex
	sync.Mutex
}

func (this *chessBoardCache) update(boardID string, f func(*models.ChessBoard) error) error {

	this.cachelock.Lock()
	boardLock, ok := this.cachelock.locks[boardID]
	if !ok {
		this.cachelock.Unlock()
		return errors.New("棋局不存在")
	}
	boardLock.Lock()
	// 拿到私有锁之后就没必要继续持有公有锁了, 立即释放以免降低性能
	this.cachelock.Unlock()
	defer boardLock.Unlock()
	if err := f(this.cache[boardID]); err != nil {
		return err
	}

	return nil
}

func (this *chessBoardCache) add(board *models.ChessBoard) error {

	boardID := board.ID.Hex()
	this.cachelock.Lock()
	defer this.cachelock.Unlock()
	_, ok := this.cachelock.locks[boardID]
	if ok {
		return errors.New("棋局已经存在")
	}
	this.cachelock.locks[boardID] = new(sync.Mutex)
	this.cache[boardID] = board
	return nil
}

func (this *chessBoardCache) remove(boardID string) {

	this.cachelock.Lock()
	defer this.cachelock.Unlock()
	if boardLock, ok := this.cachelock.locks[boardID]; ok {
		boardLock.Lock()
		delete(this.cache, boardID)
		boardLock.Unlock()
		delete(this.cachelock.locks, boardID)
	}
}

func newChessBoardCache() *chessBoardCache {
	return &chessBoardCache{make(map[string]*models.ChessBoard), &chessBoardCacheLock{locks: make(map[string]*sync.Mutex)}}
}

// 更新内存中的棋局状态
func UpdateBoardCache(boardID string, f func(*models.ChessBoard) error) error {
	return boardCache.update(boardID, f)
}

// 向内存中添加棋局
func AddBoardCache(board *models.ChessBoard) error {
	return boardCache.add(board)
}

// 删除内存中的棋局
func RemoveBoardCache(boardID string) {
	boardCache.remove(boardID)
}
