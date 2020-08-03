package supervisor

import (
	"fmt"

	"github.com/nunchistudio/blacksmith/helper/errors"
)

/*
InterfaceSupervisor is the string representation for the supervisor interface.
*/
var InterfaceSupervisor = "supervisor"

/*
Supervisor is the interface used to properly run Blacksmith applications in
distributed environments. This allows strong data consistency and better infrastructure
reliability.
*/
type Supervisor interface {

	// String returns the string representation of the adapter.
	//
	// Example: "consul"
	String() string

	// Options returns the options originally passed to the Options struct. This
	// can be used to validate and override user's options if necessary.
	Options() *Options

	// Init lets you initialize the Supervisor. It is useful to create a session
	// across nodes and register a service instance in the service registry.
	Init(*Toolkit, *Service) error

	// Shutdown let you gracefully shutdown a service in the Supervisor. It is
	// useful to destroy a session and deregister a service instance from the
	// service registry.
	Shutdown(*Toolkit, *Service) error

	// Lock allows to acquire a key in the semaphore. It returns true if the key
	// is successfully acquired.
	Lock(*Toolkit, *Lock) (bool, error)

	// Unlock allows to release a key from the semaphore. It returns true if the
	// key is successfully released.
	Unlock(*Toolkit, *Lock) (bool, error)

	// Health returns the status of the service registry.
	Health(*Toolkit) (*Registry, error)
}

/*
validateSupervisor makes sure the supervisor adapter is ready to be used properly
by a Blacksmith application.
*/
func validateSupervisor(s Supervisor) error {

	// Create the common error for the validation.
	fail := &errors.Error{
		Message:     "supervisor: Failed to load adapter",
		Validations: []errors.Validation{},
	}

	// Verify the supervisor ID is not empty.
	if s.String() == "" {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Supervisor ID must not be empty",
			Path:    []string{"Supervisor", "unknown", "String()"},
		})

		return fail
	}

	// We now can add the adapter name to the error message.
	fail.Message = fmt.Sprintf("supervisor/%s: Failed to load adapter", s.String())

	// It is impossible to deal with nil options.
	if s.Options() == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Supervisor options must not be nil",
			Path:    []string{"Supervisor", s.String(), "Options()"},
		})

		return fail
	}

	// If the adapter didn't set a node, we can not continue.
	if s.Options().Join == nil {
		fail.Validations = append(fail.Validations, errors.Validation{
			Message: "Node must not be nil",
			Path:    []string{"Supervisor", s.String(), "Options()", "Join"},
		})

		return fail
	}

	// Avoid cycles.
	s.Options().Load = nil

	// Return the error if any occurred.
	if len(fail.Validations) > 0 {
		return fail
	}

	return nil
}