package server

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net"
	"sort"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"google.golang.org/grpc"
	pb "imran/poker/protobufs"
	"imran/poker/server/game_ring"
)

const (
	Port   = ":50051"
	dbName = "poker"
)

var (
	ErrPlayerNameExists        = fmt.Errorf("player with that name already exists")
	ErrEmptyPlayerName         = fmt.Errorf("can not create player with empty name")
	ErrInvalidPlayerCount      = fmt.Errorf("can not create game with supplied count of players")
	ErrGameNameExists          = fmt.Errorf("game with that name already exists")
	ErrEmptyGameName           = fmt.Errorf("can not create game with empty name")
	ErrTooManyPlayers          = fmt.Errorf("too many players to create game")
	ErrInvalidSlotNumber       = fmt.Errorf("slot value invalid must be between 1-8")
	ErrGameDoesntExist         = fmt.Errorf("the game requesting to be updated does not exist")
	ErrInvalidButtonAllocation = fmt.Errorf("buttons are not allocated correctly")
	ErrNoBetSet                = fmt.Errorf("no bet set for game")
)

type Server struct {
	gormDb *gorm.DB
}

type Player struct {
	gorm.Model
	Name  string
	Chips int64
	Slot  int64
	H1    string
	H2    string
}

type GamePlayers struct {
	gorm.Model
	Player int64
	Game   int64
}

type Games struct {
	gorm.Model
	Name   string
	Dealer int64
	Min    int64
	f1     string
	f2     string
	f3     string
	f4     string
	f5     string
}

func NewServer(name string) (*Server, error) {
	s := &Server{}
	err := s.setupDatabase(name)
	return s, err
}

func (s *Server) setupDatabase(name string) error {

	db, err := gorm.Open("sqlite3", fmt.Sprintf("./%s.db", name))
	if err != nil {
		return err
	}
	// Setup Players table
	if err := db.AutoMigrate(&Player{}).Error; err != nil {
		return err
	}

	if err := db.AutoMigrate(&GamePlayers{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Games{}).Error; err != nil {
		return err
	}

	s.gormDb = db
	return nil
}

func (s *Server) CreatePlayer(ctx context.Context, p *pb.Player) (*pb.Player, error) {
	if p.GetName() == "" {
		return nil, ErrEmptyPlayerName
	}

	exists, err := s.GetPlayerByName(ctx, p)

	if err != nil {
		return nil, err
	}

	if exists.GetId() != 0 {
		return nil, ErrPlayerNameExists
	}

	toCreate := &Player{Name: p.GetName(), Chips: p.GetChips()}
	if err := s.gormDb.Create(toCreate).Error; err != nil {
		return nil, err
	}

	player, err := s.GetPlayer(ctx, &pb.Player{
		Id: int64(toCreate.Model.ID),
	})
	if err != nil {
		return nil, err
	}
	return player, nil

}

func (s *Server) CreatePlayers(ctx context.Context, players *pb.Players) (*pb.Players, error) {
	out := &pb.Players{}
	for _, p := range players.Players {
		player, err := s.CreatePlayer(ctx, p)
		if err != nil {
			return nil, err
		}
		out.Players = append(out.Players, player)
	}
	return out, nil

}

func (s *Server) GetPlayer(ctx context.Context, in *pb.Player) (*pb.Player, error) {

	p := &Player{
		Model: gorm.Model{
			ID: uint(in.GetId()),
		},
	}

	if err := s.gormDb.Where("id = ?", uint(in.GetId())).First(&p).Error; err != nil {
		return nil, err

	}
	return &pb.Player{
		Id:    int64(p.ID),
		Name:  p.Name,
		Chips: int64(p.Chips),
		Slot:  int64(p.Slot),
	}, nil
}

func (s *Server) GetPlayers(ctx context.Context, players *pb.Players) (*pb.Players, error) {
	outs := []Player{}

	ids := []int64{}

	for _, n := range players.GetPlayers() {
		ids = append(ids, n.GetId())
		outs = append(outs, Player{
			Model: gorm.Model{
				ID: uint(n.GetId()),
			},
		})
	}

	s.gormDb.Where("id IN (?)", ids).Find(&outs)
	out := &pb.Players{}

	for _, inp := range outs {
		out.Players = append(out.Players, &pb.Player{
			Id:    int64(inp.ID),
			Name:  inp.Name,
			Chips: inp.Chips,
			Slot:  inp.Slot,
		})

	}
	return out, nil
}

func (s *Server) GetPlayerByName(ctx context.Context, in *pb.Player) (*pb.Player, error) {
	if in.Name == "" {
		return nil, ErrEmptyPlayerName
	}
	p := &Player{
		Name: in.GetName(),
	}

	s.gormDb.Where("name", []string{"jinzhu", "jinzhu 2"}).Find(&p)
	if err := s.gormDb.Where("name = ?", in.GetName()).Find(&p).Error; err != nil && err == gorm.ErrRecordNotFound {
		return &pb.Player{}, nil
	} else if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err

	} else {

		return &pb.Player{
			Id:    int64(p.ID),
			Name:  p.Name,
			Chips: int64(p.Chips),
		}, nil
	}
}

func (s *Server) GetPlayersByName(ctx context.Context, players *pb.Players) (*pb.Players, error) {

	outs := []Player{}

	names := []string{}

	for _, n := range players.GetPlayers() {
		names = append(names, n.GetName())
		outs = append(outs, Player{
			Name: n.GetName(),
		})
	}
	s.gormDb.Where("name IN (?)", names).Find(&outs)
	out := &pb.Players{}

	// TODO switch this to be 1 query
	for _, inp := range outs {
		out.Players = append(out.Players, &pb.Player{
			Id:    int64(inp.ID),
			Name:  inp.Name,
			Chips: inp.Chips,
		})

	}
	return out, nil

}

func (s *Server) GetGame(ctx context.Context, in *pb.Game) (*pb.Game, error) {
	g := &Games{
		Model: gorm.Model{
			ID: uint(in.GetId()),
		},
	}

	if err := s.gormDb.Where("id = (?)", in.GetId()).Find(g).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, ErrGameDoesntExist
	}

	// Hydrate players
	playersId, err := s.GetGamePlayersByGameId(ctx, &pb.Game{Id: int64(g.ID)})

	if err != nil {
		return nil, err
	}

	players, err := s.GetPlayers(ctx, playersId)
	if err != nil {
		return nil, err
	}

	game := &pb.Game{
		Id:      int64(g.ID),
		Name:    g.Name,
		Players: players,
		Dealer:  g.Dealer,
		Min:     g.Min,
		// TODO hyrdate Deck and Flop

	}

	return game, nil
}

func (s *Server) GetGameByName(ctx context.Context, in *pb.Game) (*pb.Game, error) {
	g := Games{
		Name: in.GetName(),
	}

	if err := s.gormDb.Where("name = ?", in.GetName()).Find(&g).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &pb.Game{
		Id:   int64(g.ID),
		Name: g.Name,
	}, nil
}

func (s *Server) GetGamePlayersByGameId(ctx context.Context, in *pb.Game) (*pb.Players, error) {
	gp := []GamePlayers{
		{Game: in.GetId()},
	}

	if err := s.gormDb.Where("game = ?", in.GetId()).Find(&gp).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	out := &pb.Players{}
	for _, pId := range gp {
		out.Players = append(out.Players, &pb.Player{
			Id: int64(pId.Player),
		})
	}

	return out, nil
}

// SetGamePlayers Sets the game players.
// This method is flexible so if there are existing players in the game
// it will only add the difference (If the total number of players is less than 8)
func (s *Server) SetGamePlayers(ctx context.Context, g *pb.Game) (*pb.Players, error) {

	// 1. Get existing players IDs in the game
	existingIds, err := s.GetGamePlayersByGameId(ctx, g)

	if err != nil {
		return nil, err
	}

	//2. Get Existing Player Records from the IDs
	existingPlayerRecords, err := s.GetPlayers(ctx, existingIds)
	if err != nil {
		return nil, err
	}
	// 2.a create a map of existing playerIds to the player record
	existingPlayersMap := map[int64]*pb.Player{}
	for _, p := range existingPlayerRecords.GetPlayers() {
		existingPlayersMap[p.GetId()] = p
	}


	//3. Get the players requesting to be added to the game
	fmt.Println("Players to join count", g.GetPlayers().GetPlayers())

	playersToJoinRecords, err := s.GetPlayersByName(ctx, g.GetPlayers())
	if err != nil {
		return nil, err
	}
	fmt.Println("players to join", len(playersToJoinRecords.GetPlayers()))

	// 3.a create a map of requesting playerIds to whether they should join
	playersToJoinMap := map[int64]*pb.Player{}
	for _, p := range playersToJoinRecords.GetPlayers() {
		// Player is not already on the game list
		if _, ok := existingPlayersMap[p.GetId()]; !ok {
			playersToJoinMap[p.GetId()] = p
		}
	}

	for _, shouldAdd := range playersToJoinMap {
		toCreate := &GamePlayers{Player: shouldAdd.GetId(), Game: g.GetId()}
		if err := s.gormDb.Create(toCreate).Error; err != nil {
			return nil, err
		}

	}
	players, err := s.GetGamePlayersByGameId(ctx, g)
	if err != nil {
		return nil, err
	}
	return players, err

}

// SitGamePlayers allocates players to the game slots
func (s *Server) SitGamePlayers(ctx context.Context, g *pb.Game) (*pb.Game, error) {

	// 1. Get existing players IDs in the game
	existingIds, err := s.GetGamePlayersByGameId(ctx, g)

	if err != nil {
		return nil, err
	}

	// 2. Get Existing DbPlayer Records from the IDs
	existingPlayerRecords, err := s.GetPlayers(ctx, existingIds)
	if err != nil {
		return nil, err
	}

	if len(existingPlayerRecords.GetPlayers()) > 8 || len(existingPlayerRecords.GetPlayers()) < 2 {
		return nil, ErrInvalidPlayerCount
	}

	for i, p := range existingPlayerRecords.GetPlayers() {
		p.Slot = int64(i + 1)
	}

	return g, err

}

func (s *Server) SetPlayerSlot(ctx context.Context, p *pb.Player) (*pb.Player, error) {
	if p.GetSlot() > 8 || p.GetSlot() < 1 {
		return nil, ErrInvalidSlotNumber
	}

	out := &Player{
		Model: gorm.Model{
			ID: uint(p.GetId()),
		},
	}

	if err := s.gormDb.Where("id = ?", p.GetId()).Find(out).Update(
		"slot", p.GetSlot()).Error; err != nil {
		return nil, err
	}

	player, err := s.GetPlayer(ctx, p)
	if err != nil {
		return nil, err
	}
	return player, nil

}

func (s *Server) AllocateGameSlots(ctx context.Context, g *pb.Game) (*pb.Game, error) {

	players := g.GetPlayers().GetPlayers()
	for i, p := range players {
		if i > 8 {
			return nil, ErrInvalidPlayerCount
		}
		slot := i + 1
		p.Slot = int64(slot)
		_, err := s.SetPlayerSlot(ctx, p)
		if err != nil {
			return nil, err
		}
	}

	out, err := s.GetGame(ctx, g)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (s *Server) CreateGame(ctx context.Context, g *pb.Game) (*pb.Game, error) {
	if g.GetName() == "" {
		return nil, ErrEmptyGameName
	}

	exists, err := s.GetGameByName(ctx, g)
	if err != nil {
		return nil, err
	}
	if exists != nil {
		return nil, ErrGameNameExists
	}

	toCreate := &Games{
		Name:   g.GetName(),
		Dealer: g.GetDealer(),
		Min:    g.GetMin(),
	}
	if err := s.gormDb.Create(toCreate).Error; err != nil {
		return nil, err
	}

	game, err := s.GetGameByName(ctx, g)
	if err != nil {
		return nil, err
	}
	return game, nil

}

func (s *Server) SetButtonPositions(ctx context.Context, g *pb.Game) (*pb.Game, error) {
	if g.GetName() == "" {
		return nil, ErrEmptyGameName
	}

	game, err := s.GetGame(ctx, g)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameDoesntExist
	}

	toUpdate := Games{
		Min: g.GetMin(),
		// Randomly allocate a dealer
		Dealer: int64(rand.Intn(len(game.GetPlayers().GetPlayers()))) + 1,
	}

	if err := s.gormDb.Where("id = ?", game.GetId()).Find(game).Updates(toUpdate).Error; err != nil {
		return nil, err
	}

	out, err := s.GetGame(ctx, g)

	if err != nil {
		return nil, err
	}
	return out, nil

}

func (s *Server) NextDealer(ctx context.Context, g *pb.Game) (*pb.Game, error) {
	if g.GetName() == "" {
		return nil, ErrEmptyGameName
	}

	game, err := s.GetGame(ctx, g)
	if err != nil {
		return nil, err
	}
	if game == nil {
		return nil, ErrGameDoesntExist
	}

	r, err := game_ring.NewRing(g)
	if err != nil {
		return nil, err
	}

	err = r.CurrentSmallBlind()
	if err != nil {
		return nil, err
	}

	newDealer, err := r.MarshalValue()

	toUpdate := Games{
		Min: g.GetMin(),
		// Randomly allocate a dealer
		Dealer: newDealer.GetSlot(),
	}

	if err := s.gormDb.Where("id = ?", game.GetId()).Find(game).Updates(toUpdate).Error; err != nil {
		return nil, err
	}

	out, err := s.GetGame(ctx, g)

	if err != nil {
		return nil, err
	}
	return out, nil

}

// ValidatePreGame returns an error if the game is invalid
// Invalid reasons are
//  1. Not enough, or too many players
//  2. Slots are allocated to players incorrectly
//  3. Button positions and bet is not set.
func (s *Server) ValidatePreGame(ctx context.Context, g *pb.Game) (*pb.Game, error) {
	game, err := s.GetGame(ctx, g)
	if err != nil {
		return g, err
	}
	if game == nil {
		return g, ErrGameDoesntExist
	}
	// Need to have 2-8 players to play
	if len(g.GetPlayers().GetPlayers()) < 2 || len(g.GetPlayers().GetPlayers()) > 8 {
		return g, ErrInvalidPlayerCount
	}

	// mapping of user id to slot
	slotMap := map[int64]int64{}

	// mapping of slot to userId
	userMap := map[int64]int64{}

	// get a slice of all the player slots
	slotList := []int64{}

	for _, player := range g.GetPlayers().GetPlayers() {
		// Only seats 1-8 are valid
		if player.GetSlot() < 1 || player.GetSlot() > 8 {
			return g, ErrInvalidSlotNumber
		}
		// 2 players are allocated to the same slot
		if _, ok := slotMap[player.GetId()]; ok {
			return g, ErrInvalidSlotNumber
		}
		slotMap[player.GetId()] = player.GetSlot()
		userMap[player.GetSlot()] = player.GetId()

		slotList = append(slotList, player.GetSlot())
	}

	sort.Slice(slotList, func(i, j int) bool {
		return slotList[i] < slotList[j]
	})

	for i, v := range slotList {
		if !(i == 0) {
			prev := slotList[i-1]
			if !(prev < v) {
				//The slots are not sequential, or there is a gap
				return g, ErrInvalidSlotNumber
			}
		}
	}

	if g.GetDealer() == 0 {
		return g, ErrInvalidButtonAllocation
	}

	if g.GetMin() < 1 {
		return g, ErrNoBetSet
	}

	return g, nil
}

func Run() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	serv, err := NewServer(dbName)
	if err != nil {
		log.Fatalf("failed to Start poker server: %v", err)
	}
	pb.RegisterPokerServer(s, serv)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
