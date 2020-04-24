// go doc put to good use	 
				
-------------------------------------------------------------------------------
## go doc smap.SMap	
type SMap map[V]X
    SMap represents the mapping of logic Variables to symbolic eXpressions/terms
    and intentionally mimics and extends the interface of a sync.Map.

func New() SMap
func (m SMap) Clone() SMap
func (m SMap) Delete(key V)
func (m SMap) Load(key V) (value X, ok bool)
func (m SMap) LoadOrStore(key V, value X) (actual X, loaded bool)
func (m SMap) Store(key V, value X)
func (m SMap) String() string
				
-------------------------------------------------------------------------------
## go doc -all		
package smap // import "github.com/GoLangsam/kanren/internal/µ/smap"


TYPES

type SMap map[V]X
    SMap represents the mapping of logic Variables to symbolic eXpressions/terms
    and intentionally mimics and extends the interface of a sync.Map.

func New() SMap
func (m SMap) Clone() SMap
    Clone returns a shallow copy.

func (m SMap) Delete(key V)
    Delete deletes the value for a key.

func (m SMap) Load(key V) (value X, ok bool)
    Load returns the value stored in the map for a key, or nil if no value is
    present. The ok result indicates whether value was found in the map.

func (m SMap) LoadOrStore(key V, value X) (actual X, loaded bool)
    LoadOrStore returns the existing value for the key if present. Otherwise, it
    stores and returns the given value. The loaded result is true if the value
    was loaded, false if stored.

func (m SMap) Store(key V, value X)
    Store sets the value for a key.

func (m SMap) String() string
    String returns a string of the symbolic map sorted by key.

type V = X // *sexpr.Variable
    V is an eXpression which represents a logic variable

type X = *sexpr.Expression
    X represents a symbolic expression

				
-------------------------------------------------------------------------------
