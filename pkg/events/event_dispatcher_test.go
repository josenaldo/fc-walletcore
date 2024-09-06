package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 1. Configrure os tipos TestEvent e TestEventHandler para implementar as interfaces EventInterface e
//EventHandlerInterface, respectivamente.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TestEvent struct {
	Name          string
	Payload       interface{}
	eventDateTime time.Time
}

func (e TestEvent) GetName() string {
	return e.Name
}

func (e TestEvent) GetDateTime() time.Time {
	return e.eventDateTime
}

func (e TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e TestEvent) SetPayload(payload interface{}) {
	e.Payload = payload
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface) {
	// fmt.Printf("Handling event %s with handler %s\n", event.GetName(), h.name)
}

type MockHandler struct {
	mock.Mock
}

func (h *MockHandler) Handle(event EventInterface) {
	h.Called(event)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 2. Crie a suite de teste EventDispatcherTestSuite. Uma suite de teste é uma coleção de testes que podem ser
// executados juntos. A suite de teste EventDispatcherTestSuite deve incorporar suite.Suite.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type EventDispatcherTestSuite struct {
	suite.Suite
	event      TestEvent
	event2     TestEvent
	handler    TestEventHandler
	handler2   TestEventHandler
	handler3   TestEventHandler
	dispatcher *EventDispatcher
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 3. Implemente a função TestSuiteEventDispatcher. Esta função executa a suite de teste EventDispatcherTestSuite.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestSuiteEventDispatcher(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 4. Crie a função (s *EventDispatcherTestSuite) SetupTest(). Esta função é executada antes de cada teste.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *EventDispatcherTestSuite) SetupTest() {

	s.dispatcher = NewEventDispatcher()

	s.handler = TestEventHandler{
		ID: 1,
	}
	s.handler2 = TestEventHandler{
		ID: 2,
	}
	s.handler3 = TestEventHandler{
		ID: 3,
	}

	s.event = TestEvent{
		Name:          "test",
		Payload:       "test",
		eventDateTime: time.Now(),
	}

	s.event2 = TestEvent{
		Name:          "test2",
		Payload:       "test2",
		eventDateTime: time.Now(),
	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 5. Implemente a função (s *EventDispatcherTestSuite) TearDownTest(). Esta função é executada após cada teste.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *EventDispatcherTestSuite) TearDownTest() {
	s.dispatcher.Clear()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 6. Agora implemente os testes para os métodos Register, Remove, Has, Clear e Dispatch.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (suite *EventDispatcherTestSuite) TestRegisterIsSuccessful() {
	// Arrange - Given

	// Act - When
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.NoError(err)

	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(hasEventHandler)

	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 1)
	suite.Equal(suite.dispatcher.handlers[suite.event.GetName()][0], &suite.handler)
}

func (suite *EventDispatcherTestSuite) TestRegisterTwoHandlersForSameEventIsSuccessful() {
	// Arrange - Given

	// Act - When
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	err2 := suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)

	// Assert - Then
	suite.NoError(err)
	suite.NoError(err2)

	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(hasEventHandler)

	hasEventHandler2 := suite.dispatcher.Has(suite.event.GetName(), &suite.handler2)
	suite.True(hasEventHandler2)

	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 2)

	suite.Equal(suite.dispatcher.handlers[suite.event.GetName()][0], &suite.handler)
	suite.Equal(suite.dispatcher.handlers[suite.event.GetName()][1], &suite.handler2)
}

func (suite *EventDispatcherTestSuite) TestRegisterSameHandlerTwiceForSameEventReturnsError() {
	// Arrange - Given

	// Act - When
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	err2 := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.NoError(err)
	suite.Error(err2)
	suite.EqualError(err2, ErrorHandlerAlreadyRegistered.Error())

	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.True(hasEventHandler)

	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 1)
	suite.Equal(suite.dispatcher.handlers[suite.event.GetName()][0], &suite.handler)
}

func (suite *EventDispatcherTestSuite) TestClearIsSuccessful() {
	// Arrange - Given
	// Event 1
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)
	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 1)

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)
	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 2)

	//event 2
	err = suite.dispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.NoError(err)
	suite.Len(suite.dispatcher.handlers[suite.event2.GetName()], 1)

	// Act - When
	suite.dispatcher.Clear()

	// Assert - Then
	suite.Len(suite.dispatcher.handlers, 0)
}

func (suite *EventDispatcherTestSuite) TestHasReturnsFalseIfHandlerIsNotRegistered() {
	// Arrange - Given

	// Act - When
	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.False(hasEventHandler)
}

func (suite *EventDispatcherTestSuite) TestHasReturnsTrueIfHandlerIsRegistered() {
	// Arrange - Given
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)

	// Act - When
	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.True(hasEventHandler)
}

func (suite *EventDispatcherTestSuite) TestEventDispatchIsSuccessful() {
	// Arrange - Given
	eventHandler := MockHandler{}
	eventHandler.On("Handle", suite.event).Return(nil)

	err := suite.dispatcher.Register(suite.event.GetName(), &eventHandler)
	suite.NoError(err)

	// Act - When
	err = suite.dispatcher.Dispatch(suite.event)

	// Espera para dar tempo às goroutines de serem executadas
	time.Sleep(100 * time.Millisecond)

	// Assert - Then
	suite.NoError(err)
	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertCalled(suite.T(), "Handle", suite.event)
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatchWithMultipleHandlersIsSuccessful() {
	// Arrange - Given
	eventHandler := MockHandler{}
	eventHandler2 := MockHandler{}

	eventHandler.On("Handle", suite.event).Return(nil)
	eventHandler2.On("Handle", suite.event).Return(nil)

	err := suite.dispatcher.Register(suite.event.GetName(), &eventHandler)
	suite.NoError(err)

	err = suite.dispatcher.Register(suite.event.GetName(), &eventHandler2)
	suite.NoError(err)

	// Act - When
	err = suite.dispatcher.Dispatch(suite.event)

	// Espera para dar tempo às goroutines de serem executadas
	time.Sleep(100 * time.Millisecond)

	// Assert - Then
	suite.NoError(err)
	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertCalled(suite.T(), "Handle", suite.event)
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)

	eventHandler2.AssertExpectations(suite.T())
	eventHandler2.AssertCalled(suite.T(), "Handle", suite.event)
	eventHandler2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
func (suite *EventDispatcherTestSuite) TestEventDispatchWithNoHandlersRegisteredReturnsNil() {
	// Arrange - Given

	// Act - When
	err := suite.dispatcher.Dispatch(suite.event)

	// Assert - Then
	suite.NoError(err)
}

func (suite *EventDispatcherTestSuite) TestRemoveIsSuccessful() {
	// Arrange - Given
	err := suite.dispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NoError(err)

	err = suite.dispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.NoError(err)

	// Act - When
	err = suite.dispatcher.Remove(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.NoError(err)

	hasEventHandler := suite.dispatcher.Has(suite.event.GetName(), &suite.handler)
	suite.False(hasEventHandler)

	hasEventHandler2 := suite.dispatcher.Has(suite.event.GetName(), &suite.handler2)
	suite.True(hasEventHandler2)

	suite.Len(suite.dispatcher.handlers[suite.event.GetName()], 1)
	suite.Equal(suite.dispatcher.handlers[suite.event.GetName()][0], &suite.handler2)
}

func (suite *EventDispatcherTestSuite) TestRemoveReturnNilIfRemoveUnregisteredHandler() {
	// Arrange - Given

	// Act - When
	err := suite.dispatcher.Remove(suite.event.GetName(), &suite.handler)

	// Assert - Then
	suite.NoError(err)
}
